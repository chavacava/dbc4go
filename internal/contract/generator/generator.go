//go:generate dbc4go -i $GOFILE -o $GOFILE

// Package generator provides the functions necessary for generating contracts
package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/chavacava/dbc4go/internal/contract"
	contractParser "github.com/chavacava/dbc4go/internal/contract/parser"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/ast/astutil"
)

const generatedHeader = "// Code generated by dbc4go, DO NOT EDIT.\n"

// GenerateCode produces the contract enforcing code for the given source file
func GenerateCode(input io.Reader, output io.Writer) error {
	buf, err := analyzeCode(input)
	if err != nil {
		return fmt.Errorf("unable to analyze input source code, got: %w", err)
	}

	_, err = fmt.Fprintf(output, "%s", buf.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write output source code, got: %w", err)
	}

	return nil
}

// @ensures r.Len() == 0 ==> err != nil
func analyzeCode(src io.Reader) (r bytes.Buffer, err error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable to parse code: %v", err)
	}

	astDecorator := decorator.NewDecorator(fset)
	dstFile, err := astDecorator.DecorateFile(astFile)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable to decorate AST: %v", err)
	}

	fileAnalyzer := fileAnalyzer{
		decorator:          astDecorator,
		imports:            importsContainer{},
		typeInvariantsCode: map[string][]string{},
	}
	// walk the AST with the analyzer to find contracts and generate their contracts
	ast.Walk(fileAnalyzer, astFile)

	fset, astFile, err = decorator.RestoreFile(dstFile)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable to generate new code: %v", err)
	}

	// add imports required by contracts
	for k := range fileAnalyzer.imports {
		astutil.AddImport(fset, astFile, k)
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, astFile); err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable to format the new code: %w", err)
	}

	// Add standard header for generated files
	code := buf.String()
	buf.Reset()
	buf.WriteString(generatedHeader)
	buf.WriteString(code)

	finalCode := strings.ReplaceAll(buf.String(), commentPrefix, "")
	buf.Reset()
	buf.WriteString(finalCode)

	// re-parse file to check for errors in generated code
	resultFile, err := parser.ParseFile(fset, "", &buf, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("found error in generated code, please check contracts: %w on\n%s", err, addLineNumbers(finalCode))
	}

	// format generated code
	buf.Reset()
	if err := format.Node(&buf, fset, resultFile); err != nil {
		return bytes.Buffer{}, fmt.Errorf("unable to format the generated code: %w", err)
	}

	return buf, nil
}

func addLineNumbers(text string) string {
	result := ""
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		result += fmt.Sprintf("%d\t%s\n", i+1, line)
	}
	return result
}

type importsContainer map[string]struct{}

type fileAnalyzer struct {
	decorator          *decorator.Decorator
	imports            importsContainer
	typeInvariantsCode map[string][]string
}

func (fa fileAnalyzer) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fa.rewriteFuncDecl(n)
		return nil //skip visiting function body
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			return nil // not a type declaration
		}
		if len(n.Specs) <= 0 {
			return nil // no specs in the type declaration
		}
		typeSpec, ok := (n.Specs[0]).(*ast.TypeSpec)
		if !ok {
			return nil // not a type declaration
		}

		fa.analyzeTypeContract(typeSpec.Name.Name, n.Doc)
		return nil // skip visiting the type fields
	}

	return fa
}

func (fa fileAnalyzer) analyzeTypeContract(typeName string, doc *ast.CommentGroup) {
	if doc == nil {
		return // nothing to do, the type does not have associated documentation
	}

	contractParser := contractParser.NewParser()
	contract := contract.NewTypeContract(typeName)
	for _, commentLine := range doc.List {
		err := contractParser.ParseTypeContract(contract, commentLine.Text)
		if err != nil {
			log.Printf("%s: Warning: %s", fa.positionAsString(commentLine.Pos()), err.Error())
			continue
		}
	}

	fa.addCodeForTypeInvariant(typeName, contract)
}

func (fa fileAnalyzer) addCodeForTypeInvariant(typeName string, contract *contract.TypeContract) {
	fa.typeInvariantsCode[typeName] = fa.generateInvariantCode(contract)
}

// positionAsString returns a string representation of the given token position
// @requires fa.decorator.Fset != nil
func (fa fileAnalyzer) positionAsString(pos token.Pos) string {
	position := fa.decorator.Fset.Position(pos)

	return position.Filename + ":" + strconv.Itoa(position.Line) + ":" + strconv.Itoa(position.Column)
}

// rewriteFuncDecl is in charge of generating contract-enforcing code for functions
// @requires fd != nil
func (fa *fileAnalyzer) rewriteFuncDecl(fd *ast.FuncDecl) {
	contract.OldCounter = 0 // TODO: refactor
	dstFuncDecl := fa.decorator.Dst.Nodes[fd].(*dst.FuncDecl)
	if fd.Doc != nil {
		contractParser := contractParser.NewParser()
		contract := contract.NewFuncContract(fd)
		comments := fd.Doc.List
		for _, commentLine := range comments {
			err := contractParser.ParseFuncContract(contract, commentLine.Text)
			if err != nil {
				log.Printf("%s: Warning: %s", fa.positionAsString(commentLine.Pos()), err.Error())
				continue
			}
		}

		contractStmts, errs := fa.generateCode(contract)
		for _, err := range errs {
			log.Printf("Warning: %v", err)
		}

		dstFuncDecl.Body.Decorations().Start.Append(contractStmts...)
	}

	// Also add code for enforce invariants if available
	if fd.Recv == nil || len(fd.Recv.List) < 1 {
		return // not a method thus no invariants
	}

	receiverType := fa.getReceiverTypeName(fd.Recv)
	invariantCode, ok := fa.typeInvariantsCode[receiverType]
	if !ok || invariantCode == nil {
		return // did not found invariant code associated to this method's receiver
	}
	if len(fd.Recv.List[0].Names) < 1 || fd.Recv.List[0].Names[0].Name == "_" {
		// anonymous receiver
		log.Printf("Warning: can not enforce invariants on method %s because it has an anonymous receiver", fd.Name.Name)
		return
		// TODO: insert a receiver name to enable checks
	}

	receiverName := fd.Recv.List[0].Names[0].Name
	invariantCodeForMethod := make([]string, len(invariantCode))
	for i, code := range invariantCode {
		invariantCodeForMethod[i] = strings.ReplaceAll(code, receiverType+".", receiverName+".")
	}
	dstFuncDecl.Body.Decorations().Start.Append(invariantCodeForMethod...)
}

func (fa fileAnalyzer) getReceiverTypeName(receiver *ast.FieldList) string {
	if len(receiver.List) < 1 {
		return "UNKNOWN"
	}

	recType := receiver.List[0].Type
	return strings.Replace(fa.typeAsString(recType), "*", "", 1)
}

// generateCode yields the list of GO statements that enforce the given contract
// It also yields the list of errors that occurred while the generation
// @requires c != nil
// All ensures are grouped in a single defer statement and there is an if statement for each require
// @ensures len(c.Ensures()) > 0 ==> len(c.Requires())+1 == len(stmts)+len(errs)
// @ensures len(c.Ensures()) == 0 ==> len(c.Requires()) == len(stmts)+len(errs)
func (fa fileAnalyzer) generateCode(c *contract.FuncContract) (stmts []string, errs []error) {
	result := []string{}
	errs = []error{}
	for _, r := range c.Requires() {
		stmt := fa.generateRequiresCode(r)
		result = append(result, stmt)
	}

	if len(c.Ensures()) > 0 {
		stmt := fa.generateEnsuresCode(c.Ensures(), c.Target())
		result = append(result, stmt)
	}

	// merge new imports into imports list
	for k, v := range c.Imports() {
		fa.imports[k] = v
	}

	return result, errs
}

const commentPrefix = "//dbc4go "

// @requires c != nil
// @ensures len(c.Ensures()) == 0 ==> stmts == nil
// @ensures len(c.Ensures()) != 0 ==> stmts != nil
func (fa fileAnalyzer) generateInvariantCode(c *contract.TypeContract) (stmts []string) {
	if len(c.Ensures()) == 0 {
		return nil
	}
	result := []string{}
	const templateEnsure = commentPrefix + `if %shortStmt%!(%cond%) { panic("(type invariant) %contract% not satisfied") }`
	clauses := c.Ensures()
	ensuresCode := make([]string, len(clauses))
	for _, clause := range clauses {
		shortStmt, expr, _ := clause.ExpandedExpression()
		if shortStmt != "" {
			shortStmt = shortStmt + "; "
		}
		ensure := strings.Replace(templateEnsure, "%shortStmt%", shortStmt, 1)
		ensure = strings.Replace(ensure, "%cond%", expr, 1)
		ensure = strings.Replace(ensure, "%contract%", escapeDoubleQuotes(clause.String()), 1)
		ensuresCode = append(ensuresCode, ensure)
	}
	const templateDeferredFunction = commentPrefix + `defer func(){%checks%}()`
	r := strings.Replace(templateDeferredFunction, "%checks%", strings.Join(ensuresCode, "\n"), 1)
	result = append(result, r)

	// merge new imports into imports list
	for k, v := range c.Imports() {
		fa.imports[k] = v
	}

	return result
}

func (fileAnalyzer) generateRequiresCode(req contract.Requires) (r string) {
	const templateRequire = commentPrefix + `if !(%cond%) { panic("%contract% not satisfied") }`
	exp := req.ExpandedExpression()

	r = strings.Replace(templateRequire, "%cond%", exp, 1)
	r = strings.Replace(r, "%contract%", escapeDoubleQuotes(req.String()), 1)

	return r
}

// @requires fd != nil
// @requires clauses != nil && len(clauses) > 0
// @ensures r != ""
func (fa fileAnalyzer) generateEnsuresCode(clauses []contract.Ensures, fd *ast.FuncDecl) (r string) {
	const templateEnsure = commentPrefix + `if %shortStmt%!(%cond%) { panic("%contract% not satisfied") }`

	ensuresCode := make([]string, len(clauses))
	funcParams := []string{}
	funcArgs := []string{}
	for _, clause := range clauses {
		shortStmt, expr, idToOld := clause.ExpandedExpression()
		if shortStmt != "" {
			shortStmt = shortStmt + "; "
		}
		for id, old := range idToOld {
			funcParams = append(funcParams, fmt.Sprintf("%s %s", old, fa.getTypeForID(id, fd)))
			funcArgs = append(funcArgs, id)
		}
		ensure := strings.Replace(templateEnsure, "%shortStmt%", shortStmt, 1)
		ensure = strings.Replace(ensure, "%cond%", expr, 1)
		ensure = strings.Replace(ensure, "%contract%", escapeDoubleQuotes(clause.String()), 1)
		ensuresCode = append(ensuresCode, ensure)
	}

	const templateDeferredFunction = commentPrefix + `defer func(%params%){%checks%}(%args%)`

	r = strings.Replace(templateDeferredFunction, "%params%", strings.Join(funcParams, ","), 1)
	r = strings.Replace(r, "%checks%", strings.Join(ensuresCode, "\n"), 1)
	r = strings.Replace(r, "%args%", strings.Join(funcArgs, ","), 1)

	return r
}

// @requires fd != nil
func (fa fileAnalyzer) getTypeForID(id string, fd *ast.FuncDecl) string {
	for _, param := range fd.Type.Params.List {
		fields := param.Names
		found := false
		for _, field := range fields {
			name := field.Name
			if id == name {
				found = true
				break
			}
		}

		if found {
			return fa.typeAsString(param.Type)
		}
	}

	return "any"
}

// @requires n != nil
// @ensures r != ""
func (fa fileAnalyzer) typeAsString(n ast.Node) (r string) {
	buf := bytes.Buffer{}
	fs := token.NewFileSet()
	printer.Fprint(&buf, fs, n)
	return buf.String()
}

// @ensures len(str)<=len(r)
func escapeDoubleQuotes(str string) (r string) {
	return strings.Replace(str, "\"", "\\\"", -1)
}

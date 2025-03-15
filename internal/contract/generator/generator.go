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

// GenerateCode produces the contract enforcing code for the given source file.
//
// Contract:
//   - requires input != nil
//   - requires output != nil
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

// Contract:
//   - requires src != nil
//   - ensures err != nil ==> r.Len() == 0
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

// Contract:
//   - requires typeName != ""
func (fa fileAnalyzer) analyzeTypeContract(typeName string, doc *ast.CommentGroup) {
	if doc == nil {
		return // nothing to do, the type does not have associated documentation
	}

	parser := contractParser.NewParser()
	contract, err := parser.ParseTypeContract(typeName, doc.List)
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	if contract.IsEmpty() {
		return // found no contracts for this type
	}

	fa.addCodeForTypeInvariant(typeName, contract)
}

// Contract:
//   - requires typeName != ""
//   - requires contract != nil
//   - requires fa.typeInvariantsCode != nil
func (fa fileAnalyzer) addCodeForTypeInvariant(typeName string, contract *contract.TypeContract) {
	fa.typeInvariantsCode[typeName] = fa.generateInvariantCode(contract)
}

// positionAsString returns a string representation of the given token position.
//
// Contract:
//   - requires fa.decorator.Fset != nil
func (fa fileAnalyzer) positionAsString(pos token.Pos) string {
	position := fa.decorator.Fset.Position(pos)

	return position.Filename + ":" + strconv.Itoa(position.Line) + ":" + strconv.Itoa(position.Column)
}

// rewriteFuncDecl is in charge of generating contract-enforcing code for functions.
//
// Contract:
//   - requires fd != nil
func (fa *fileAnalyzer) rewriteFuncDecl(fd *ast.FuncDecl) {
	contract.OldCounter = 0 // TODO: refactor
	dstFuncDecl := fa.decorator.Dst.Nodes[fd].(*dst.FuncDecl)

	contractScope := []string{commentPrefix + "{ // Open contract scope "}

	defer func() {
		codeWasGenerated := len(contractScope) > 1
		if codeWasGenerated {
			contractScope = append(contractScope, commentPrefix+"} // Close contract scope")
			dstFuncDecl.Body.Decorations().Start.Append(contractScope...)
		}
	}()

	// Add code for enforce invariants if any
	contractScope = append(contractScope, fa.getCodeForInvariants(fd)...)

	// Add code for function contracts if any
	contractStmts, errs := fa.getCodeForContracts(fd)
	for _, err := range errs {
		log.Printf("Warning: %v", err)
	}

	contractScope = append(contractScope, contractStmts...)
}

// Contract:
//   - requires fd != nil
//   - ensures fd.Doc == nil ==> len(result) == 0 && len(errs) == 0
func (fa fileAnalyzer) getCodeForContracts(fd *ast.FuncDecl) (result []string, errs []error) {
	if fd.Doc == nil {
		return result, errs // the function has not attached documentation
	}

	parser := contractParser.NewParser()
	contract, err := parser.ParseFuncContract(fd.Doc.List)
	if err != nil {
		return result, append(errs, err)
	}

	if contract.IsEmpty() {
		return result, errs // the function has not attached contract
	}

	result, errs = fa.generateCode(contract)
	if len(result) > 0 {
		resultWithComment := []string{commentPrefix + "// Function's contracts"}
		result = append(resultWithComment, result...)
	}

	return result, errs
}

// Contract:
//   - requires fd != nil
func (fa fileAnalyzer) getCodeForInvariants(fd *ast.FuncDecl) (result []string) {
	if fd.Recv == nil || len(fd.Recv.List) < 1 {
		return result
	}

	receiverType := fa.getReceiverTypeName(fd.Recv)
	invariantCode, ok := fa.typeInvariantsCode[receiverType]
	if !ok || len(invariantCode) == 0 {
		return result
	}

	isAnonymousReceiver := len(fd.Recv.List[0].Names) < 1 || fd.Recv.List[0].Names[0].Name == "_"
	if isAnonymousReceiver {
		log.Printf("Warning: can not enforce invariants on method %s because it has an anonymous receiver", fd.Name.Name)
		return result
		// TODO: insert a receiver name to enable checks
	}

	receiverName := fd.Recv.List[0].Names[0].Name
	invariantCodeForMethod := make([]string, len(invariantCode))
	for i, code := range invariantCode {
		invariantCodeForMethod[i] = strings.ReplaceAll(code, receiverType+".", receiverName+".")
	}

	result = []string{commentPrefix + "// Type invariants "}
	result = append(result, invariantCodeForMethod...)

	return result
}

// Contract:
//   - requires receiver != nil
func (fa fileAnalyzer) getReceiverTypeName(receiver *ast.FieldList) string {
	if len(receiver.List) < 1 {
		return "UNKNOWN"
	}

	recType := receiver.List[0].Type
	return strings.Replace(fa.typeAsString(recType), "*", "", 1)
}

// generateCode yields the list of GO statements that enforce the given contract.
// It also yields the list of errors that occurred while the generation.
//
// Contract:
//   - requires contract != nil
//   - unmodified len(contract.Requires())
//   - unmodified len(contract.Ensures())
//   - unmodified len(contract.Lets())
//   - unmodified len(contract.Imports())
//   - unmodified contract.Target()
func (fa fileAnalyzer) generateCode(contract *contract.FuncContract) (stmts []string, errs []error) {
	result := []string{}
	errs = []error{}

	for _, r := range contract.Requires() {
		stmt := fa.generateRequiresCode(r, "")
		result = append(result, stmt)
	}

	for _, let := range contract.Lets() {
		stmt := fa.generateLetCode(let)
		result = append(result, stmt)
	}

	if len(contract.Ensures()) > 0 {
		stmt := fa.generateEnsuresCode(contract.Ensures())
		result = append(result, stmt)
	}

	// merge new imports into imports list
	for k, v := range contract.Imports() {
		fa.imports[k] = v
	}

	return result, errs
}

const commentPrefix = "//dbc4go "

// Contract:
//   - requires c != nil
//   - ensures len(c.Ensures()) == 0 ==> len(stmts) == 0
//   - ensures len(c.Ensures()) != 0 ==> len(stmts) > 0
func (fa fileAnalyzer) generateInvariantCode(c *contract.TypeContract) (stmts []string) {
	result := []string{}

	// Generate requires for invariants
	for _, req := range c.Requires() {
		result = append(result, fa.generateRequiresCode(req, "(type invariant) "))
	}

	// Generate ensures for invariants
	if len(c.Ensures()) != 0 {
		const templateEnsure = commentPrefix + `if %shortStmt%!(%cond%) { panic("(type invariant) function didn't ensure %contract%") }`
		clauses := c.Ensures()
		ensuresCode := make([]string, len(clauses))
		for _, clause := range clauses {
			shortStmt, expr, _ := clause.ExpandedExpression()
			if shortStmt != "" {
				shortStmt = shortStmt + "; "
			}
			ensure := strings.Replace(templateEnsure, "%shortStmt%", shortStmt, 1)
			ensure = strings.Replace(ensure, "%cond%", expr, 1)
			contractStr := clause.Description()
			if contractStr == "" {
				contractStr = escapeDoubleQuotes(clause.Expression().Raw)
			}
			ensure = strings.Replace(ensure, "%contract%", contractStr, 1)
			ensuresCode = append(ensuresCode, ensure)
		}
		const templateDeferredFunction = commentPrefix + `defer func(){%checks%}()`
		r := strings.Replace(templateDeferredFunction, "%checks%", strings.Join(ensuresCode, "\n"), 1)
		result = append(result, r)
	}

	// merge new imports into imports list
	for k, v := range c.Imports() {
		fa.imports[k] = v
	}

	return result
}

func (fileAnalyzer) generateRequiresCode(req contract.Requires, panicMsgPrefix string) (r string) {
	const templateRequire = commentPrefix + `if !(%cond%) { panic("%msgPrefix%function caller didn't satisfied %contract%") }`
	exp := req.ExpandedExpression()

	contractStr := req.Description()
	if contractStr == "" {
		contractStr = escapeDoubleQuotes(req.Expression().Raw)
	}
	r = strings.Replace(templateRequire, "%cond%", exp.Raw, 1)
	r = strings.Replace(r, "%msgPrefix%", panicMsgPrefix, 1)
	r = strings.Replace(r, "%contract%", contractStr, 1)

	return r
}

func (fileAnalyzer) generateLetCode(let contract.Let) (r string) {
	const templateLet = commentPrefix + `%decl% // %description%`
	r = strings.Replace(templateLet, `%decl%`, let.ExpandedExpression().Raw, 1)
	description := let.Description()
	if description == "" {
		description = " defined with @let"
	}
	r = strings.Replace(r, "%description%", description, 1)

	return r
}

func (fa fileAnalyzer) generateEnsuresCode(clauses []contract.Ensures) (r string) {
	ensuresCode := make([]string, len(clauses))
	oldVarDecls := []string{}
	for i, clause := range clauses {
		code, decls := generateEnsuresCodeFromExpression(clause.Expression(), clause.Description(), true)
		oldVarDecls = append(oldVarDecls, decls...)
		ensuresCode[i] = code
	}

	r += strings.Join(oldVarDecls, "\n")
	if len(r) > 0 {
		r += "\n"
	}

	const templateDeferredFunction = commentPrefix + "defer func(){\n%checks%}()"
	r += strings.Replace(templateDeferredFunction, "%checks%", strings.Join(ensuresCode, "\n"), 1)

	return r
}

func generateEnsuresCodeFromExpression(expression contract.Expression, description string, isRootExpression bool) (code string, oldVarDecls []string) {
	switch expression.Kind {
	case contract.ExprKindPlain:
		return generateEnsuresCodeFromPlainExpression(expression, description, isRootExpression)
	case contract.ExprKindForall:
		return generateEnsuresCodeFromForallExpression(expression, description, isRootExpression)
	case contract.ExprKindExists:
		return generateEnsuresCodeFromExistsExpression(expression, description, isRootExpression)
	default:
		log.Panicf("Unknown expression kind %d", expression.Kind)
	}
	return
}

func generateEnsuresCodeFromPlainExpression(expression contract.Expression, description string, isRootExpression bool) (code string, oldVarDecls []string) {
	const templateOldVarDecl = commentPrefix + `%oldId% := %expr%`
	templateEnsure := templateRootPlainExpression
	if !isRootExpression {
		templateEnsure = strings.ReplaceAll(templateLeafPlainExpression, "\n", " ")
	}
	templateEnsure = commentPrefix + templateEnsure

	shortStmt, expr, idToOld := contract.ExpandEnsuresExpression(expression)
	if shortStmt != "" {
		shortStmt = shortStmt + "; "
	}

	for expr, oldID := range idToOld {
		decl := strings.Replace(templateOldVarDecl, "%oldId%", oldID, 1)
		decl = strings.Replace(decl, "%expr%", expr, 1)
		oldVarDecls = append(oldVarDecls, decl)
	}

	code = strings.Replace(templateEnsure, "%shortStmt%", shortStmt, 1)
	code = strings.Replace(code, "%cond%", expr, 1)
	contractStr := description
	if contractStr == "" {
		contractStr = escapeDoubleQuotes(expression.Raw)
	}

	code = strings.Replace(code, "%contract%", contractStr, 1)
	return code, oldVarDecls
}

func generateEnsuresCodeFromForallExpression(expression contract.Expression, description string, isRootExpression bool) (code string, oldVarDecls []string) {
	//const templateForallElement = commentPrefix + `for %variable% := range %source% { %expression% }`
	//const templateForallIndex = commentPrefix + `for %variable%:=0;%variable%<len(%source%);%variable%++ { %expression% }`

	variable := expression.SubExprs[contract.ExprKindForallFieldVariables]
	source := expression.SubExprs[contract.ExprKindForallFieldSources]
	if description == "" {
		description = expression.Raw
	}

	rangeOrIteration := ""
	forallKind := expression.SubExprs[contract.ExprKindForallFieldKind].Raw
	switch forallKind {
	case contract.ForallKindIn:
		rangeOrIteration = "%variable% := range %source%"
	case contract.ForallKindIndexof:
		rangeOrIteration = "%variable%:=0;%variable%<len(%source%);%variable%++"
	default:
		log.Panicf("Unknown @forall kind %s", forallKind)
	}

	subExpressionCode, varDecls := generateEnsuresCodeFromExpression(expression.SubExprs[contract.ExprKindForallFieldExpression], description, false)
	template := commentPrefix + strings.ReplaceAll(templateForallExpression, "\n", " ")
	code = strings.ReplaceAll(template, "%rangeOrIteration%", rangeOrIteration)
	code = strings.ReplaceAll(code, "%variable%", variable.Raw)
	code = strings.ReplaceAll(code, "%source%", source.Raw)
	code = strings.Replace(code, "%subexpression%", subExpressionCode, 1)

	action := `panic("function didn't ensure %contract%")`
	alternativeAction := ""
	if !isRootExpression {
		action = "return false"
		alternativeAction = "return true"
	}
	code = strings.Replace(code, "%action%", action, 1)
	code = strings.Replace(code, "%alternativeAction%", alternativeAction, 1)
	code = strings.Replace(code, "%contract%", expression.Raw, 1)

	return code, varDecls
}

func generateEnsuresCodeFromExistsExpression(expression contract.Expression, description string, isRootExpression bool) (code string, oldVarDecls []string) {
	variable := expression.SubExprs[contract.ExprKindExistsFieldVariables]
	source := expression.SubExprs[contract.ExprKindExistsFieldSources]
	if description == "" {
		description = expression.Raw
	}

	rangeOrIteration := ""
	existKind := expression.SubExprs[contract.ExprKindExistsFieldKind].Raw
	switch existKind {
	case contract.ForallKindIn:
		rangeOrIteration = "%variable% := range %source%"
	case contract.ForallKindIndexof:
		rangeOrIteration = "%variable%:=0;%variable%<len(%source%);%variable%++"
	default:
		log.Panicf("Unknown @exist kind %s", existKind)
	}

	subExpressionCode, varDecls := generateEnsuresCodeFromExpression(expression.SubExprs[contract.ExprKindForallFieldExpression], description, false)
	template := commentPrefix + strings.ReplaceAll(templateExistsExpression, "\n", " ")
	code = strings.ReplaceAll(template, "%rangeOrIteration%", rangeOrIteration)
	code = strings.ReplaceAll(code, "%variable%", variable.Raw)
	code = strings.ReplaceAll(code, "%source%", source.Raw)
	code = strings.Replace(code, "%subexpression%", subExpressionCode, 1)

	action := `panic("function didn't ensure %contract%")`
	alternativeAction := ""
	if !isRootExpression {
		action = "return false"
		alternativeAction = "return true"
	}
	code = strings.Replace(code, "%action%", action, 1)
	code = strings.Replace(code, "%alternativeAction%", alternativeAction, 1)
	code = strings.Replace(code, "%contract%", expression.Raw, 1)

	return code, varDecls
}

// Contract:
//   - requires n != nil
//   - ensures r != ""
func (fa fileAnalyzer) typeAsString(n ast.Node) (r string) {
	buf := bytes.Buffer{}
	fs := token.NewFileSet()
	printer.Fprint(&buf, fs, n)
	return buf.String()
}

// Contract:
//   - ensures len(str)<=len(r)
func escapeDoubleQuotes(str string) (r string) {
	return strings.Replace(str, "\"", "\\\"", -1)
}

const templateRootPlainExpression = `if %shortStmt%!(%cond%) { panic("function didn't ensure %contract%") }`
const templateLeafPlainExpression = `
return %cond%`

const templateForallExpression = `
cond:= func() bool {
	for %rangeOrIteration% {
		cond := func() bool {
			%subexpression%
		};
		if !cond() {
			return false
		}		
	};
	return true
};
if !cond() {
	%action%
};
%alternativeAction%`

const templateExistsExpression = `
cond:= func() bool {
	for %rangeOrIteration% {
		cond := func() bool {
			%subexpression%
		};
		if cond() {
			return true
		}		
	};
	return false
};
if !cond() {
	%action%
};
%alternativeAction%`

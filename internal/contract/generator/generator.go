//go:generate dbc4go -i $GOFILE -o $GOFILE
package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chavacava/dbc4go/internal/astutils"
	"github.com/chavacava/dbc4go/internal/contract"
	cparser "github.com/chavacava/dbc4go/internal/contract/parser"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/fatih/astrewrite"
)

const generatedHeader = "// Code generated by dbc4go, DO NOT EDIT.\n"

// generateCode yields the list of GO statements that enforce the given contract
// It also yields the list of errors that occurred while the generation
//@requires c != nil
//@ensures len(c.Ensures()) > 0 ==> len(c.Requires())+1 == len(stmts)+len(errs)
//@ensures len(c.Ensures()) == 0 ==> len(c.Requires()) == len(stmts)+len(errs)
func generateCode(c *contract.FuncContract) (stmts []ast.Stmt, imports map[string]struct{}, errs []error) {
	result := []ast.Stmt{}
	errs = []error{}
	for _, r := range c.Requires() {
		stmt, err := generateRequiresCode(r)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to generate code for the clause '%s'", r))
			continue
		}

		result = append(result, stmt)
	}

	if len(c.Ensures()) > 0 {
		stmt, err := generateEnsuresCode(c.Ensures(), c.Target())
		if err != nil {
			errs = append(errs, errors.Wrap(err, "unable to generate code for @ensure clause"))
		} else {
			result = append(result, stmt)
		}
	}

	return result, c.Imports(), errs
}

// must ensure r == nil => e != nil
//@ensures r != nil || e != nil
func generateRequiresCode(req contract.Requires) (r ast.Stmt, e error) {
	exp := req.ExpandedExpression()
	expAST, err := parser.ParseExpr("!(" + exp + ")")
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse expression '%s'", exp)
	}

	msgAST := astutils.NewStringLit("\"precondition " + escapeDoubleQuotes(exp) + " not satisfied\"")
	panicArgs := astutils.NewCallArgs(msgAST)
	call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
	body := astutils.NewStmtBlock(call2panic)

	return astutils.NewIf(expAST, *body), nil
}

//@requires fd != nil
//@requires len(clauses) > 0
//@ensure r == nil ==> e != nil
func generateEnsuresCode(clauses []contract.Ensures, fd *ast.FuncDecl) (r ast.Stmt, e error) {
	funcBody := []ast.Stmt{}

	for _, clause := range clauses {
		exp := clause.ExpandedExpression()
		expAST, err := parser.ParseExpr("!(" + exp + ")")
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse expression '%s'", exp)
		}

		msgAST := astutils.NewStringLit("\"postcondition " + escapeDoubleQuotes(exp) + " not satisfied\"")
		panicArgs := astutils.NewCallArgs(msgAST)
		call2panic := astutils.NewCallAsStmt("", "panic", panicArgs)
		body := astutils.NewStmtBlock(call2panic)
		funcBody = append(funcBody, astutils.NewIf(expAST, *body))
	}

	funcParams := []*ast.Field{}
	funcArgs := []ast.Expr{}
	if fd.Type.Params != nil {
		funcParams = append(funcParams, astutils.CopyFields(fd.Type.Params.List, "old_")...)
		funcArgs = append(funcArgs, astutils.ArgsFromFields(fd.Type.Params.List)...)
	}
	if fd.Recv != nil {
		funcParams = append(funcParams, astutils.CopyFields(fd.Recv.List, "old_")...)
		funcArgs = append(funcArgs, astutils.ArgsFromFields(fd.Recv.List)...)
	}

	funcCall := astutils.NewCallAnonymous(funcParams, astutils.NewStmtBlock(funcBody...), funcArgs)

	return astutils.NewDeferStmt(funcCall), nil
}

//@ensures len(str)<=len(r)
func escapeDoubleQuotes(str string) (r string) {
	return strings.Replace(str, "\"", "\\\"", -1)
}

// GenerateCode produces the contract enforcing code for the given source file
//@requires inputFilename != "" && outputFilename != ""
func GenerateCode(inputFilename, outputFilename string) error {
	// simplify analyzeCode by just setting inputFilename
	inputSourceCode, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		return errors.Wrap(err, "unable to open input file")
	}
	buf, err := analyzeCode(inputSourceCode, inputFilename)
	if err != nil {
		return errors.Wrap(err, "unable to analyze input source code")
	}

	output := os.Stdout

	if outputFilename != "" {
		var err error
		output, err = os.Create(outputFilename)
		if err != nil {
			return errors.Wrapf(err, "unable to create output file '%s'", outputFilename)
		}
		defer output.Close()

		log.Printf("Generating file '%s'", outputFilename)
	}

	_, err = fmt.Fprintf(output, "%s", buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "unable to write output source code")
	}

	return nil
}

func analyzeCode(src []byte, fileName string) (bytes.Buffer, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, src, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not parse code: %v", err)
	}

	fr := fileRewriter{fset: fset, file: file}
	newFile := astrewrite.Walk(file, fr.rewrite)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, newFile); err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not format the new code: %v", err)
	}

	if fr.modified {
		// Add standard header for generated files
		code := buf.String()
		buf.Reset()
		buf.WriteString(generatedHeader)
		buf.WriteString(code)
	}

	return buf, nil
}

type fileRewriter struct {
	fset     *token.FileSet
	file     *ast.File
	modified bool
}

// rewrite rewrites an AST node
// this function is to be used with astrewrite.Walk
func (fr *fileRewriter) rewrite(node ast.Node) (ast.Node, bool) {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fd, imports := fr.rewriteFuncDecl(n)
		fSet := token.NewFileSet()
		for k := range imports {
			astutil.AddImport(fSet, fr.file, k)
		}
		return fd, true
	}

	return node, true
}

// positionAsString returns a string representation of the given token position
//@requires fr.fset != nil
func (fr *fileRewriter) positionAsString(pos token.Pos) string {
	position := fr.fset.Position(pos)
	return position.Filename + ":" + strconv.Itoa(position.Line) + ":" + strconv.Itoa(position.Column)
}

// rewriteFuncDecl is in charge of generating contract-enforcing code for functions
//@requires fd != nil
//@ensures r != nil
func (fr *fileRewriter) rewriteFuncDecl(fd *ast.FuncDecl) (r *ast.FuncDecl, imports map[string]struct{}) {
	if fd.Doc == nil {
		return fd, map[string]struct{}{} // nothing to do, the function does not have a comment
	}

	cp := cparser.NewParser()
	contract := contract.NewFuncContract(fd)
	comments := fd.Doc.List
	for _, commentLine := range comments {
		err := cp.Parse(&contract, commentLine.Text)
		if err != nil {
			log.Printf("%s: Warning: %s", fr.positionAsString(commentLine.Pos()), err.Error())
			continue
		}
	}

	contractStmts, imports, errs := generateCode(&contract)
	for _, err := range errs {
		log.Printf("Warning: %v", err)
	}

	fd.Body.List = append(contractStmts, fd.Body.List...)
	fr.modified = fr.modified || len(contractStmts) > 0

	return fd, imports
}

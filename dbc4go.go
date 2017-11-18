package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"strings"

	"github.com/chavacava/dbc4go/astutils"
)

const (
	requiresNonil = iota
	requiresExp   = iota
)

func main() {
	srcFile := flag.String("i", "", "input source file")
	flag.Parse()

	src, err := ioutil.ReadFile(*srcFile)
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}

	buf, err := analyzeCode(src, *srcFile)

	if err != nil {
		log.Fatalf("could not analyze source code: %v", err)
	}

	fmt.Printf("%s", buf.Bytes())
}

func analyzeCode(src []byte, fileName string) (bytes.Buffer, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, src, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not parse input code: %v", err)
	}

	var genError error
	ast.Inspect(file, func(x ast.Node) bool {
		s, ok := x.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if contracts := getContractComments(s.Doc.Text()); len(contracts) > 0 {
			log.Printf("Function %s has %d contracts\n", s.Name, len(contracts))
			for _, c := range contracts {
				err := generateContractCode(c, s)
				if err != nil {
					genError = err
					return true
				}
			}
		}
		return false
	})

	if genError != nil {
		return bytes.Buffer{}, genError
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not format the new code: %v", err)
	}

	return buf, nil
}

func getContractComments(c string) []string {
	result := []string{}
	for _, line := range strings.Split(strings.TrimSuffix(c, "\n"), "\n") {
		if strings.HasPrefix(line, "@") {
			result = append(result, line)
		}
	}
	return result
}

func generateContractCode(c string, f *ast.FuncDecl) error {
	ctype, err := getContractType(c)
	if err != nil {
		return fmt.Errorf("could not generate contract code for '%s': %v", c, err)
	}

	switch ctype {
	case requiresNonil:
		checks := []ast.Stmt{}
		for _, id := range getFuncParams(f) {
			checks = append(checks, getNilCheckingStmt(id))
		}

		f.Body.List = append(checks, f.Body.List...)
	case requiresExp:
		exp := c[len("@\\requires"):]
		check, err := getExpCheckingStmt(exp)
		if err != nil {
			return err
		}
		f.Body.List = append([]ast.Stmt{check}, f.Body.List...)
	default:
		return fmt.Errorf("could not generate contract code for '%s': unknown contract type %d", c, ctype)
	}

	return nil
}

func getContractType(c string) (int, error) {
	if strings.HasPrefix(c, "@\\requires nonil") {
		return requiresNonil, nil
	}
	if strings.HasPrefix(c, "@\\requires") {
		return requiresExp, nil
	}

	return -1, fmt.Errorf("could not identify the contrat for '%s'", c)
}

// getNilCheckingStmt yields a stmt that checks if the given id is nil.
func getNilCheckingStmt(id string) ast.Stmt {
	args := astutils.NewCallArgs(astutils.NewStringLit("\"value of parameter '" + id + "' must not be nil\""))
	logCall := astutils.NewCallAsStmt("log", "Fatal", args)
	cond := astutils.NewBinExp(token.EQL, astutils.NewID(id), astutils.NewID("nil"))
	body := astutils.NewStmtBlock(logCall)

	return astutils.NewIf(cond, *body)
	// you can do astutils.NewIf(cond, *body, astutils.WithElse(elseBody))
}

func getExpCheckingStmt(exp string) (ast.Stmt, error) {
	expr, err := parser.ParseExpr("!(" + exp + ")")
	if err != nil {
		log.Printf("error in expression: %v", err)
		return nil, fmt.Errorf("error while parsing expression %s: %v", exp, err)
	}

	args := astutils.NewCallArgs(astutils.NewStringLit("\"condition (" + exp + " ) must be met\""))
	logCall := astutils.NewCallAsStmt("log", "Fatal", args)
	body := astutils.NewStmtBlock(logCall)

	return astutils.NewIf(expr, *body), nil
}

func getFuncParams(fd *ast.FuncDecl) []string {
	result := []string{}

	for _, field := range fd.Type.Params.List {

		for _, n := range field.Names {
			result = append(result, n.Name)
		}
	}

	return result
}

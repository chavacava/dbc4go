//go:generate dbc4go -i $GOFILE -o $GOFILE
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
	"os"
	"strconv"

	"github.com/chavacava/dbc4go/internal/contract/generator"
	cparser "github.com/chavacava/dbc4go/internal/contract/parser"
	cast "github.com/chavacava/dbc4go/internal/contract/parser/ast"

	"github.com/fatih/astrewrite"
)

func main() {
	srcFile := flag.String("i", "", "input source file")
	targetFile := flag.String("o", "", "ouput file (defaults to stdout")
	flag.Parse()

	if *srcFile == "" {
		log.Fatal("Undefined input file, please set the flag -i")
	}

	src, err := ioutil.ReadFile(*srcFile)
	if err != nil {
		log.Fatalf("Could not open input file: %v", err)
	}

	buf, err := analyzeCode(src, *srcFile)

	if err != nil {
		log.Fatalf("Could not analyze source code: %v", err)
	}

	target := os.Stdout

	if *targetFile != "" {
		var err error
		target, err = os.Create(*targetFile)
		if err != nil {
			log.Fatalf("Unable to create output file '%s': %v", *targetFile, err)
		}
		defer target.Close()
		log.Printf("Generating file '%s'", *targetFile)
	}

	fmt.Fprintf(target, "%s", buf.Bytes())
}

func analyzeCode(src []byte, fileName string) (bytes.Buffer, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, src, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not parse code: %v", err)
	}

	fr := fileRewriter{fset: fset}
	newFile := astrewrite.Walk(file, fr.rewrite)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, newFile); err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not format the new code: %v", err)
	}

	return buf, nil
}

type fileRewriter struct {
	fset *token.FileSet
}

// rewrite rewrites an AST node
// this function is to be used with astrewrite.Walk
func (fr *fileRewriter) rewrite(node ast.Node) (ast.Node, bool) {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fd := fr.rewriteFuncDecl(n)
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
func (fr *fileRewriter) rewriteFuncDecl(fd *ast.FuncDecl) *ast.FuncDecl {
	if fd.Doc == nil {
		return fd // nothing to do, the function does not have a comment
	}

	cp := cparser.NewParser()
	contract := cast.NewContract(fd)
	comments := fd.Doc.List
	for _, commentLine := range comments {
		err := cp.Parse(&contract, commentLine.Text)
		if err != nil {
			log.Printf("%s: Warning: %s", fr.positionAsString(commentLine.Pos()), err.Error())
			continue
		}

		contractStmts, errs := generator.GenerateCode(&contract)
		for _, err := range errs {
			log.Printf("%s: Warning: %v", fr.positionAsString(commentLine.Pos()), err)
		}

		fd.Body.List = append(contractStmts, fd.Body.List...)
	}

	return fd
}

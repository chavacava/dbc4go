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
	"strconv"

	"github.com/chavacava/dbc4go/contract"
	"github.com/fatih/astrewrite"
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
	_ = buf
}

func analyzeCode(src []byte, fileName string) (bytes.Buffer, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, src, parser.ParseComments)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("could not parse input code: %v", err)
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
		fd, cont := fr.rewriteFuncDecl(n)
		return fd, cont
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
func (fr *fileRewriter) rewriteFuncDecl(fd *ast.FuncDecl) (*ast.FuncDecl, bool) {
	if fd.Doc == nil {
		return fd, true // nothing to do, the function does not have a comment
	}

	cp := contract.Parser{}
	comments := fd.Doc.List
	for _, commentLine := range comments {

		checkCodeRoot, err := cp.Parse(commentLine.Text)

		if err != nil {
			log.Printf("%s: Warning: %s", fr.positionAsString(commentLine.Pos()), err.Error())
			continue
		}

		if checkCodeRoot == nil && err == nil {
			continue // not a comment containing a contract
		}

		fd.Body.List = append([]ast.Stmt{checkCodeRoot}, fd.Body.List...)
	}

	return fd, true
}

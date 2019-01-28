package generator

import (
	"bytes"
	"go/ast"
	"strings"
	"testing"

	"github.com/chavacava/dbc4go/internal/astutils"
	"github.com/chavacava/dbc4go/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrivateGenerateCode(t *testing.T) {
	tests := []struct {
		requires []string
		ensures  []string
		imports  []string
	}{
		{
			requires: []string{},
			ensures:  []string{},
			imports:  []string{},
		},
		{
			requires: []string{"a>b"},
			ensures:  []string{},
			imports:  []string{},
		},
		{
			requires: []string{"a>b", "b == c"},
			ensures:  []string{"r != nil"},
			imports:  []string{},
		},
		{
			requires: []string{},
			ensures:  []string{"r != nil", "r == nil"},
			imports:  []string{"paths", "package"},
		},
	}

	// func NewFuncDecl(name string, formalParams *ast.FieldList, body *ast.BlockStmt) *ast.FuncDecl {

	fd := astutils.NewFuncDecl("foo", &ast.FieldList{}, &ast.BlockStmt{})

	for _, tc := range tests {
		c := contract.NewFuncContract(fd)

		for _, r := range tc.requires {
			c.AddRequires(contract.NewRequires(r))
		}

		for _, e := range tc.ensures {
			c.AddEnsures(contract.NewEnsures(e))
		}

		for _, i := range tc.imports {
			c.AddImport(i)
		}

		stmts, imports, errs := generateCode(&c)

		if len(tc.ensures) == 0 {
			assert.Len(t, stmts, len(tc.requires))
		} else {
			assert.Len(t, stmts, len(tc.requires)+1)
		}

		assert.Len(t, imports, len(tc.imports))
		assert.Empty(t, errs)
	}
}

func TestGenerateCode(t *testing.T) {
	in := strings.NewReader(input)
	out := bytes.NewBufferString("")

	err := GenerateCode(in, out)

	require.Nil(t, err)
	assert.Equal(t, out.String(), output)
}

var input = `package foo

import "fmt"

//@requires a > 0 ==> b > 0
//@ensures a==0 ==> r==0
//@import "strings"
func bar(a int, b int) (r int) {
    return a
} 

//@ensures @old(a)==0 ==> r==0
//@import strings
func bar2(a int) (r int) {
    return a
}`

var output = `// Code generated by dbc4go, DO NOT EDIT.
package foo

import (
	"fmt"
	"strings"
)

//@requires a > 0 ==> b > 0
//@ensures a==0 ==> r==0
//@import "strings"
func bar(a int, b int) (r int) {
	if !(!(a > 0) ||

		(b > 0)) {
		panic("precondition !(a > 0) || (b > 0) not satisfied")
	}
	defer func(old_a int, old_b int) {
		if !(!(a == 0) ||

			(r == 0)) {
			panic("postcondition !(a==0) || (r==0) not satisfied")
		}
	}(a, b)
	return a
}

//@ensures @old(a)==0 ==> r==0
//@import strings
func bar2(a int) (r int) {
	defer func(old_a int) {
		if !(!(old_a == 0) ||
			(r == 0)) {
			panic("postcondition !(old_a==0) || (r==0) not satisfied")
		}
	}(a)
	return a
}
`
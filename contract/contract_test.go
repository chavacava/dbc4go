package contract

import (
	"github.com/chavacava/dbc4go/astutils"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"testing"
)

func requiresTemplater(cond string) string {
	return "if !(" + cond + `) {
		panic("precondition a == a not satisfied")
	}`
}

func TestGenerateRequires(t *testing.T) {

	var tt = []struct {
		prop     string
		expected string
	}{
		{"a == a",
			`if !(a == a) {
	panic("precondition a == a not satisfied")
}`},
		{"a > b && c",
			`if !(a > b && c) {
	panic("precondition a > b && c not satisfied")
}`},
		{"true",
			`if !(true) {
	panic("precondition true not satisfied")
}`},
	}

	inputGenerator := func(expr string) ast.Expr {
		expAST, err := parser.ParseExpr("!(" + expr + ")")
		if err != nil {
			return nil
		}

		return expAST
	}

	cp := Parser{}
	for _, test := range tt {
		if r := astutils.GOfmt(cp.generateRequires(inputGenerator(test.prop), test.prop)); r != test.expected {
			t.Errorf("for %s expected\n%s\n\nbut got \n\n%s", test.prop, test.expected, r)
		}
	}
}

func TestParseRequires(t *testing.T) {
	tt := []struct {
		prop string
		code string
		err  string
	}{
		{
			"qsd qsd",
			"",
			"unable to parse @require proposition 'qsd qsd': 1:7: expected ')', found qsd",
		},
		{
			"a == b",
			`if !(a == b) {
	panic("precondition a == b not satisfied")
}`,
			""},
	}

	cp := Parser{}
	for _, test := range tt {
		n, err := cp.parseRequires(test.prop)

		assert.Equal(t, astutils.GOfmt(n), test.code)
		if test.err != "" {
			assert.EqualError(t, err, test.err)
		} else {
			assert.NoError(t, err)
		}

	}
}

func TestParse(t *testing.T) {
	tt := []struct {
		line string
		code string
		err  string // "" means no error
	}{
		{
			"// This is a comment",
			"",
			"",
		},
		{
			"//@badcontract an expression",
			"",
			"unknown contract kind badcontract",
		},
		{
			"//@requires sdsd)1",
			"",
			"unable to parse @require proposition 'sdsd)1': 1:8: expected 'EOF', found 1",
		},
		{
			"//@requires a == b",
			`if !(a == b) {
	panic("precondition a == b not satisfied")
}`,
			""},
	}

	cp := Parser{}
	for _, test := range tt {
		n, err := cp.Parse(test.line)

		assert.Equal(t, astutils.GOfmt(n), test.code)
		if test.err != "" {
			assert.EqualError(t, err, test.err)
		} else {
			assert.NoError(t, err)
		}
	}
}

package contract

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chavacava/dbc4go/internal/contract/generator/parser"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func TestGenerator(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "true",
			want:  "true",
		},
		{
			input: "!true",
			want:  "!true",
		},
		{
			input: "true == false",
			want:  "true == false",
		},
		{
			input: "!true == false",
			want:  "!true == false",
		},
		{
			input: "(true != false)",
			want:  "(true != false)",
		},
		{
			input: "true ==> false",
			want:  "implies(true,false)",
		},
		{
			input: "@forall e @in elements: someCondition",
			want:  "foralle(e,elements,someCondition)",
		},
		{
			input: "@forall e @in elements: someCondition ==> false",
			want:  "foralle(e,elements,implies(someCondition,false))",
		},
		{
			input: "@forall e @in elements: someCondition ==> @forall e2 @in elements2: a+b == 5",
			want:  "foralle(e,elements,implies(someCondition,foralle(e2,elements2,a+b==5)))",
		},
		{
			input: "@forall i @indexof elements: elements[i] == false",
			want:  "foralli(i,elements,elements[i] == false)",
		},
		{
			input: "@exists i @indexof elements: elements[i] == false",
			want:  "existsi(i,elements,elements[i] == false)",
		},
		{
			input: "@exists e @in elements: e == call(something+1,slice[i])",
			want:  "existse(e,elements,e==call(something+1,slice[i]))",
		},
	}

	for _, tc := range testCases {
		// Setup the input
		is := antlr.NewInputStream(tc.input)

		// Create the Lexer
		lexer := parser.NewClauseExpressionLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewClauseExpressionParser(stream)

		// Finally parse the clause
		listener := &contractGenerator{stack: *lls.New()}
		antlr.ParseTreeWalkerDefault.Walk(listener, p.ClauseExpression())

		got, _ := listener.stack.Pop()
		want := strings.ReplaceAll(tc.want, " ", "")
		if want != got {
			t.Fatalf("Expected:\n\t%s\ngot:\n\t%s", want, got)
		}
	}
}

package contract

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chavacava/dbc4go/internal/generator/parser"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func TestGenerator(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "true",
			want:  "cond := func() bool {return true};",
		},
		{
			input: "!true",
			want:  "cond := func() bool {return !true};",
		},
		{
			input: "true == false",
			want:  "cond := func() bool {return true==false};",
		},
		{
			input: "!true == false",
			want:  "cond := func() bool {return !true==false};",
		},
		{
			input: "(true != false)",
			want:  "cond := func() bool {return (true!=false)};",
		},
		{
			input: "true ==> false",
			want:  "cond:= func() bool {cond1 := func() bool {return true};;cond2 := func() bool {return false};;return !cond1() || cond2()};",
		},
		{
			input: "@forall e @in elements: someCondition",
			want:  "cond := func() bool {for _,e := range elements {cond := func() bool {return someCondition};if !cond() {return false}}; return true};",
		},
		{
			input: "@forall e @in elements: someCondition ==> false",
			want:  "cond := func() bool {for _,e := range elements {cond:= func() bool {cond1 := func() bool {return someCondition};;cond2 := func() bool {return false};;return !cond1() || cond2()};if !cond() {return false}}; return true};",
		},
		{
			input: "@forall e @in elements: someCondition ==> @forall e2 @in elements2: a+b == 5",
			want:  "cond := func() bool {for _,e := range elements {cond:= func() bool {cond1 := func() bool {return someCondition};;cond2 := func() bool {for _,e2 := range elements2 {cond := func() bool {return a+b==5};if !cond() {return false}}; return true};;return !cond1() || cond2()};if !cond() {return false}}; return true};",
		},
		{
			input: "@forall i @indexof elements: elements[i] == false",
			want:  "cond := func() bool {for i,_ := range elements {cond := func() bool {return elements[i]==false};if !cond() {return false}}; return true};",
		},
		{
			input: "@exists i @indexof elements: elements[i] == false",
			want:  "cond := func() bool {for i,_ := range elements {cond := func() bool {return elements[i]==false};if cond() {return true}}; return false};",
		},
		{
			input: "@exists e @in elements: e == call(something+1,slice[i])",
			want:  "cond := func() bool {for _,e := range elements {cond := func() bool {return e==call(something+1,slice[i])};if cond() {return true}}; return false};",
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
		if tc.want != got {
			t.Fatalf("Expected:\n\t%s\ngot:\n\t%s", tc.want, got)
		}
	}
}

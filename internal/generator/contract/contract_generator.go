package contract

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chavacava/dbc4go/internal/generator/parser"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

// contractGenerator allows generating a clause contract by listenent to events
// from the clause expression parser
type contractGenerator struct {
	parser.BaseClauseExpressionListener
	stack lls.Stack
}

func (g *contractGenerator) ExitCompleteGoExpression(ctx *parser.CompleteGoExpressionContext) {
	cond := fmt.Sprintf("cond := func() bool {return %s};", ctx.GetText())
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitIterator(ctx *parser.IteratorContext) {
	g.stack.Push(ctx.GetText())
}

func (g *contractGenerator) ExitCollection(ctx *parser.CollectionContext) {
	g.stack.Push(ctx.GetText())
}

// ExitImplies is called when production Implies is exited.
func (g *contractGenerator) ExitImplies(ctx *parser.ImpliesContext) {
	right, _ := g.stack.Pop()
	left, _ := g.stack.Pop()
	rightR := strings.Replace(right.(string), "cond", "cond2", 1)
	leftR := strings.Replace(left.(string), "cond", "cond1", 1)
	expr := fmt.Sprintf("cond:= func() bool {%s;%s;return !cond1() || cond2()};", leftR, rightR)
	g.stack.Push(expr)
}

func (g *contractGenerator) ExitIff(ctx *parser.IffContext) {
	right, _ := g.stack.Pop()
	left, _ := g.stack.Pop()
	rightR := strings.Replace(right.(string), "cond", "cond2", 1)
	leftR := strings.Replace(left.(string), "cond", "cond1", 1)
	expr := fmt.Sprintf("cond:= func() bool {%s;%s;return (!cond1() || cond2()) && (!cond2() || cond1())};", leftR, rightR)
	g.stack.Push(expr)
}

// ExitForallElement is called when production ForallElement is exited.
func (g *contractGenerator) ExitForallElement(ctx *parser.ForallElementContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	element, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for _,%s := range %s {%sif !cond() {return false}}; return true};", element, collection, expr)
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitForallIndex(ctx *parser.ForallIndexContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	index, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for %s,_ := range %s {%sif !cond() {return false}}; return true};", index, collection, expr)
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitForallIterator(ctx *parser.ForallIteratorContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	index, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for %s := range %s {%sif !cond() {return false}}; return true};", index, collection, expr)
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitExistsElement(ctx *parser.ExistsElementContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	element, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for _,%s := range %s {%sif cond() {return true}}; return false};", element, collection, expr)
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitExistsIndex(ctx *parser.ExistsIndexContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	index, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for %s,_ := range %s {%sif cond() {return true}}; return false};", index, collection, expr)
	g.stack.Push(cond)
}

func (g *contractGenerator) ExitExistsIterator(ctx *parser.ExistsIteratorContext) {
	expr, _ := g.stack.Pop()
	collection, _ := g.stack.Pop()
	index, _ := g.stack.Pop()
	cond := fmt.Sprintf("cond := func() bool {for %s := range %s {%sif cond() {return true}}; return false};", index, collection, expr)
	g.stack.Push(cond)
}

func Generate(input string) string {
	// Setup the input
	is := antlr.NewInputStream(input)

	g := &contractGenerator{stack: *lls.New()}

	// Create the Lexer
	lexer := parser.NewClauseExpressionLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewClauseExpressionParser(stream)

	// Finally parse the clause
	antlr.ParseTreeWalkerDefault.Walk(g, p.ClauseExpression())

	result, _ := g.stack.Pop()
	return (result).(string)
}

// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ClauseExpression

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ClauseExpressionParser struct {
	*antlr.BaseParser
}

var ClauseExpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clauseexpressionParserInit() {
	staticData := &ClauseExpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'!'", "'('", "')'", "'=='", "'!='", "'||'", "'&&'", "'>'",
		"'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "','", "'['",
		"']'", "'@forall'", "'@exists'", "'==>'", "'@in'", "'@indexof'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "FORALL", "EXISTS", "IMPLIES", "IN", "INDEXOF", "DOT",
		"WHITESPACE", "ID", "DECIMAL_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
	}
	staticData.RuleNames = []string{
		"root", "clauseExpression", "iterator", "collection", "completeGoExpression",
		"goExpression", "primaryExpression", "qualifiedIdentifier", "functionCallArguments",
		"sliceIndex", "number", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 133, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 59, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 65, 8, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 70, 8, 1, 10, 1, 12, 1, 73, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 84, 8, 5, 10, 5, 12, 5, 87, 9, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 100,
		8, 6, 3, 6, 102, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 107, 8, 7, 10, 7, 12, 7,
		110, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 116, 8, 8, 10, 8, 12, 8, 119,
		9, 8, 3, 8, 121, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 0, 1, 2, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 0, 2, 1, 0, 5, 17, 1, 0, 30, 31, 137, 0, 24, 1, 0, 0, 0, 2, 64, 1,
		0, 0, 0, 4, 74, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8, 78, 1, 0, 0, 0, 10, 80,
		1, 0, 0, 0, 12, 101, 1, 0, 0, 0, 14, 103, 1, 0, 0, 0, 16, 111, 1, 0, 0,
		0, 18, 124, 1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 130, 1, 0, 0, 0, 24, 25,
		3, 2, 1, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0,
		28, 29, 5, 21, 0, 0, 29, 30, 3, 4, 2, 0, 30, 31, 5, 24, 0, 0, 31, 32, 3,
		6, 3, 0, 32, 33, 5, 1, 0, 0, 33, 34, 3, 2, 1, 6, 34, 65, 1, 0, 0, 0, 35,
		36, 5, 21, 0, 0, 36, 37, 3, 4, 2, 0, 37, 38, 5, 25, 0, 0, 38, 39, 3, 6,
		3, 0, 39, 40, 5, 1, 0, 0, 40, 41, 3, 2, 1, 5, 41, 65, 1, 0, 0, 0, 42, 43,
		5, 22, 0, 0, 43, 44, 3, 4, 2, 0, 44, 45, 5, 24, 0, 0, 45, 46, 3, 6, 3,
		0, 46, 47, 5, 1, 0, 0, 47, 48, 3, 2, 1, 4, 48, 65, 1, 0, 0, 0, 49, 50,
		5, 22, 0, 0, 50, 51, 3, 4, 2, 0, 51, 52, 5, 25, 0, 0, 52, 53, 3, 6, 3,
		0, 53, 54, 5, 1, 0, 0, 54, 55, 3, 2, 1, 3, 55, 65, 1, 0, 0, 0, 56, 65,
		3, 8, 4, 0, 57, 59, 5, 2, 0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 61, 5, 3, 0, 0, 61, 62, 3, 2, 1, 0, 62, 63, 5,
		4, 0, 0, 63, 65, 1, 0, 0, 0, 64, 27, 1, 0, 0, 0, 64, 35, 1, 0, 0, 0, 64,
		42, 1, 0, 0, 0, 64, 49, 1, 0, 0, 0, 64, 56, 1, 0, 0, 0, 64, 58, 1, 0, 0,
		0, 65, 71, 1, 0, 0, 0, 66, 67, 10, 7, 0, 0, 67, 68, 5, 23, 0, 0, 68, 70,
		3, 2, 1, 8, 69, 66, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 3, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 28,
		0, 0, 75, 5, 1, 0, 0, 0, 76, 77, 5, 28, 0, 0, 77, 7, 1, 0, 0, 0, 78, 79,
		3, 10, 5, 0, 79, 9, 1, 0, 0, 0, 80, 85, 3, 12, 6, 0, 81, 82, 7, 0, 0, 0,
		82, 84, 3, 12, 6, 0, 83, 81, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 11, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88,
		89, 5, 2, 0, 0, 89, 102, 3, 12, 6, 0, 90, 91, 5, 3, 0, 0, 91, 92, 3, 10,
		5, 0, 92, 93, 5, 4, 0, 0, 93, 102, 1, 0, 0, 0, 94, 102, 3, 20, 10, 0, 95,
		102, 3, 22, 11, 0, 96, 99, 3, 14, 7, 0, 97, 100, 3, 16, 8, 0, 98, 100,
		3, 18, 9, 0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 102, 1, 0, 0, 0, 101, 88, 1, 0, 0, 0, 101, 90, 1, 0, 0, 0, 101,
		94, 1, 0, 0, 0, 101, 95, 1, 0, 0, 0, 101, 96, 1, 0, 0, 0, 102, 13, 1, 0,
		0, 0, 103, 108, 5, 28, 0, 0, 104, 105, 5, 26, 0, 0, 105, 107, 5, 28, 0,
		0, 106, 104, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 15, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 120, 5,
		3, 0, 0, 112, 117, 3, 10, 5, 0, 113, 114, 5, 18, 0, 0, 114, 116, 3, 10,
		5, 0, 115, 113, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0,
		117, 118, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120,
		112, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123,
		5, 4, 0, 0, 123, 17, 1, 0, 0, 0, 124, 125, 5, 19, 0, 0, 125, 126, 3, 10,
		5, 0, 126, 127, 5, 20, 0, 0, 127, 19, 1, 0, 0, 0, 128, 129, 5, 29, 0, 0,
		129, 21, 1, 0, 0, 0, 130, 131, 7, 1, 0, 0, 131, 23, 1, 0, 0, 0, 9, 58,
		64, 71, 85, 99, 101, 108, 117, 120,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ClauseExpressionParserInit initializes any static state used to implement ClauseExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClauseExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClauseExpressionParserInit() {
	staticData := &ClauseExpressionParserStaticData
	staticData.once.Do(clauseexpressionParserInit)
}

// NewClauseExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClauseExpressionParser(input antlr.TokenStream) *ClauseExpressionParser {
	ClauseExpressionParserInit()
	this := new(ClauseExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ClauseExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ClauseExpression.g4"

	return this
}

// ClauseExpressionParser tokens.
const (
	ClauseExpressionParserEOF                    = antlr.TokenEOF
	ClauseExpressionParserT__0                   = 1
	ClauseExpressionParserT__1                   = 2
	ClauseExpressionParserT__2                   = 3
	ClauseExpressionParserT__3                   = 4
	ClauseExpressionParserT__4                   = 5
	ClauseExpressionParserT__5                   = 6
	ClauseExpressionParserT__6                   = 7
	ClauseExpressionParserT__7                   = 8
	ClauseExpressionParserT__8                   = 9
	ClauseExpressionParserT__9                   = 10
	ClauseExpressionParserT__10                  = 11
	ClauseExpressionParserT__11                  = 12
	ClauseExpressionParserT__12                  = 13
	ClauseExpressionParserT__13                  = 14
	ClauseExpressionParserT__14                  = 15
	ClauseExpressionParserT__15                  = 16
	ClauseExpressionParserT__16                  = 17
	ClauseExpressionParserT__17                  = 18
	ClauseExpressionParserT__18                  = 19
	ClauseExpressionParserT__19                  = 20
	ClauseExpressionParserFORALL                 = 21
	ClauseExpressionParserEXISTS                 = 22
	ClauseExpressionParserIMPLIES                = 23
	ClauseExpressionParserIN                     = 24
	ClauseExpressionParserINDEXOF                = 25
	ClauseExpressionParserDOT                    = 26
	ClauseExpressionParserWHITESPACE             = 27
	ClauseExpressionParserID                     = 28
	ClauseExpressionParserDECIMAL_LIT            = 29
	ClauseExpressionParserRAW_STRING_LIT         = 30
	ClauseExpressionParserINTERPRETED_STRING_LIT = 31
)

// ClauseExpressionParser rules.
const (
	ClauseExpressionParserRULE_root                  = 0
	ClauseExpressionParserRULE_clauseExpression      = 1
	ClauseExpressionParserRULE_iterator              = 2
	ClauseExpressionParserRULE_collection            = 3
	ClauseExpressionParserRULE_completeGoExpression  = 4
	ClauseExpressionParserRULE_goExpression          = 5
	ClauseExpressionParserRULE_primaryExpression     = 6
	ClauseExpressionParserRULE_qualifiedIdentifier   = 7
	ClauseExpressionParserRULE_functionCallArguments = 8
	ClauseExpressionParserRULE_sliceIndex            = 9
	ClauseExpressionParserRULE_number                = 10
	ClauseExpressionParserRULE_string                = 11
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ClauseExpression() IClauseExpressionContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClauseExpressionParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.clauseExpression(0)
	}
	{
		p.SetState(25)
		p.Match(ClauseExpressionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClauseExpressionContext is an interface to support dynamic dispatch.
type IClauseExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsClauseExpressionContext differentiates from other interfaces.
	IsClauseExpressionContext()
}

type ClauseExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseExpressionContext() *ClauseExpressionContext {
	var p = new(ClauseExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_clauseExpression
	return p
}

func InitEmptyClauseExpressionContext(p *ClauseExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_clauseExpression
}

func (*ClauseExpressionContext) IsClauseExpressionContext() {}

func NewClauseExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseExpressionContext {
	var p = new(ClauseExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_clauseExpression

	return p
}

func (s *ClauseExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseExpressionContext) CopyAll(ctx *ClauseExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ClauseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprInParensContext struct {
	ClauseExpressionContext
}

func NewExprInParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprInParensContext {
	var p = new(ExprInParensContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ExprInParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprInParensContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ExprInParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterExprInParens(s)
	}
}

func (s *ExprInParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitExprInParens(s)
	}
}

func (s *ExprInParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitExprInParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImpliesContext struct {
	ClauseExpressionContext
}

func NewImpliesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImpliesContext {
	var p = new(ImpliesContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ImpliesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpliesContext) AllClauseExpression() []IClauseExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			len++
		}
	}

	tst := make([]IClauseExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClauseExpressionContext); ok {
			tst[i] = t.(IClauseExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ImpliesContext) ClauseExpression(i int) IClauseExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ImpliesContext) IMPLIES() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserIMPLIES, 0)
}

func (s *ImpliesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterImplies(s)
	}
}

func (s *ImpliesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitImplies(s)
	}
}

func (s *ImpliesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitImplies(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistsIndexContext struct {
	ClauseExpressionContext
}

func NewExistsIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistsIndexContext {
	var p = new(ExistsIndexContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ExistsIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsIndexContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserEXISTS, 0)
}

func (s *ExistsIndexContext) Iterator() IIteratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorContext)
}

func (s *ExistsIndexContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserINDEXOF, 0)
}

func (s *ExistsIndexContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ExistsIndexContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ExistsIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterExistsIndex(s)
	}
}

func (s *ExistsIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitExistsIndex(s)
	}
}

func (s *ExistsIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitExistsIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistsElementContext struct {
	ClauseExpressionContext
}

func NewExistsElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistsElementContext {
	var p = new(ExistsElementContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ExistsElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsElementContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserEXISTS, 0)
}

func (s *ExistsElementContext) Iterator() IIteratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorContext)
}

func (s *ExistsElementContext) IN() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserIN, 0)
}

func (s *ExistsElementContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ExistsElementContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ExistsElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterExistsElement(s)
	}
}

func (s *ExistsElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitExistsElement(s)
	}
}

func (s *ExistsElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitExistsElement(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlainGoExpressionContext struct {
	ClauseExpressionContext
}

func NewPlainGoExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainGoExpressionContext {
	var p = new(PlainGoExpressionContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *PlainGoExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainGoExpressionContext) CompleteGoExpression() ICompleteGoExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompleteGoExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompleteGoExpressionContext)
}

func (s *PlainGoExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterPlainGoExpression(s)
	}
}

func (s *PlainGoExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitPlainGoExpression(s)
	}
}

func (s *PlainGoExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitPlainGoExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForallIndexContext struct {
	ClauseExpressionContext
}

func NewForallIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForallIndexContext {
	var p = new(ForallIndexContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ForallIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForallIndexContext) FORALL() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserFORALL, 0)
}

func (s *ForallIndexContext) Iterator() IIteratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorContext)
}

func (s *ForallIndexContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserINDEXOF, 0)
}

func (s *ForallIndexContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ForallIndexContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ForallIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterForallIndex(s)
	}
}

func (s *ForallIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitForallIndex(s)
	}
}

func (s *ForallIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitForallIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForallElementContext struct {
	ClauseExpressionContext
}

func NewForallElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForallElementContext {
	var p = new(ForallElementContext)

	InitEmptyClauseExpressionContext(&p.ClauseExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClauseExpressionContext))

	return p
}

func (s *ForallElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForallElementContext) FORALL() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserFORALL, 0)
}

func (s *ForallElementContext) Iterator() IIteratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorContext)
}

func (s *ForallElementContext) IN() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserIN, 0)
}

func (s *ForallElementContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ForallElementContext) ClauseExpression() IClauseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseExpressionContext)
}

func (s *ForallElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterForallElement(s)
	}
}

func (s *ForallElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitForallElement(s)
	}
}

func (s *ForallElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitForallElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) ClauseExpression() (localctx IClauseExpressionContext) {
	return p.clauseExpression(0)
}

func (p *ClauseExpressionParser) clauseExpression(_p int) (localctx IClauseExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewClauseExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IClauseExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ClauseExpressionParserRULE_clauseExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForallElementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(28)
			p.Match(ClauseExpressionParserFORALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Iterator()
		}
		{
			p.SetState(30)
			p.Match(ClauseExpressionParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Collection()
		}
		{
			p.SetState(32)
			p.Match(ClauseExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.clauseExpression(6)
		}

	case 2:
		localctx = NewForallIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(ClauseExpressionParserFORALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Iterator()
		}
		{
			p.SetState(37)
			p.Match(ClauseExpressionParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Collection()
		}
		{
			p.SetState(39)
			p.Match(ClauseExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.clauseExpression(5)
		}

	case 3:
		localctx = NewExistsElementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(ClauseExpressionParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Iterator()
		}
		{
			p.SetState(44)
			p.Match(ClauseExpressionParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Collection()
		}
		{
			p.SetState(46)
			p.Match(ClauseExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.clauseExpression(4)
		}

	case 4:
		localctx = NewExistsIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(ClauseExpressionParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Iterator()
		}
		{
			p.SetState(51)
			p.Match(ClauseExpressionParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Collection()
		}
		{
			p.SetState(53)
			p.Match(ClauseExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.clauseExpression(3)
		}

	case 5:
		localctx = NewPlainGoExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.CompleteGoExpression()
		}

	case 6:
		localctx = NewExprInParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ClauseExpressionParserT__1 {
			{
				p.SetState(57)
				p.Match(ClauseExpressionParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(60)
			p.Match(ClauseExpressionParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.clauseExpression(0)
		}
		{
			p.SetState(62)
			p.Match(ClauseExpressionParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImpliesContext(p, NewClauseExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ClauseExpressionParserRULE_clauseExpression)
			p.SetState(66)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				goto errorExit
			}
			{
				p.SetState(67)
				p.Match(ClauseExpressionParserIMPLIES)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.clauseExpression(8)
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIteratorContext is an interface to support dynamic dispatch.
type IIteratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIteratorContext differentiates from other interfaces.
	IsIteratorContext()
}

type IteratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteratorContext() *IteratorContext {
	var p = new(IteratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_iterator
	return p
}

func InitEmptyIteratorContext(p *IteratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_iterator
}

func (*IteratorContext) IsIteratorContext() {}

func NewIteratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteratorContext {
	var p = new(IteratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_iterator

	return p
}

func (s *IteratorContext) GetParser() antlr.Parser { return s.parser }

func (s *IteratorContext) ID() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserID, 0)
}

func (s *IteratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IteratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterIterator(s)
	}
}

func (s *IteratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitIterator(s)
	}
}

func (s *IteratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitIterator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) Iterator() (localctx IIteratorContext) {
	localctx = NewIteratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ClauseExpressionParserRULE_iterator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(ClauseExpressionParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) ID() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserID, 0)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (s *CollectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitCollection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ClauseExpressionParserRULE_collection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(ClauseExpressionParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompleteGoExpressionContext is an interface to support dynamic dispatch.
type ICompleteGoExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GoExpression() IGoExpressionContext

	// IsCompleteGoExpressionContext differentiates from other interfaces.
	IsCompleteGoExpressionContext()
}

type CompleteGoExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompleteGoExpressionContext() *CompleteGoExpressionContext {
	var p = new(CompleteGoExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_completeGoExpression
	return p
}

func InitEmptyCompleteGoExpressionContext(p *CompleteGoExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_completeGoExpression
}

func (*CompleteGoExpressionContext) IsCompleteGoExpressionContext() {}

func NewCompleteGoExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompleteGoExpressionContext {
	var p = new(CompleteGoExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_completeGoExpression

	return p
}

func (s *CompleteGoExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CompleteGoExpressionContext) GoExpression() IGoExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoExpressionContext)
}

func (s *CompleteGoExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompleteGoExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompleteGoExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterCompleteGoExpression(s)
	}
}

func (s *CompleteGoExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitCompleteGoExpression(s)
	}
}

func (s *CompleteGoExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitCompleteGoExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) CompleteGoExpression() (localctx ICompleteGoExpressionContext) {
	localctx = NewCompleteGoExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ClauseExpressionParserRULE_completeGoExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.GoExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoExpressionContext is an interface to support dynamic dispatch.
type IGoExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrimaryExpression() []IPrimaryExpressionContext
	PrimaryExpression(i int) IPrimaryExpressionContext

	// IsGoExpressionContext differentiates from other interfaces.
	IsGoExpressionContext()
}

type GoExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoExpressionContext() *GoExpressionContext {
	var p = new(GoExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_goExpression
	return p
}

func InitEmptyGoExpressionContext(p *GoExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_goExpression
}

func (*GoExpressionContext) IsGoExpressionContext() {}

func NewGoExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoExpressionContext {
	var p = new(GoExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_goExpression

	return p
}

func (s *GoExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoExpressionContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GoExpressionContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *GoExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterGoExpression(s)
	}
}

func (s *GoExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitGoExpression(s)
	}
}

func (s *GoExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitGoExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) GoExpression() (localctx IGoExpressionContext) {
	localctx = NewGoExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ClauseExpressionParserRULE_goExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.PrimaryExpression()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(81)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&262112) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(82)
				p.PrimaryExpression()
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	GoExpression() IGoExpressionContext
	Number() INumberContext
	String_() IStringContext
	QualifiedIdentifier() IQualifiedIdentifierContext
	FunctionCallArguments() IFunctionCallArgumentsContext
	SliceIndex() ISliceIndexContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionContext) GoExpression() IGoExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoExpressionContext)
}

func (s *PrimaryExpressionContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PrimaryExpressionContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *PrimaryExpressionContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *PrimaryExpressionContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *PrimaryExpressionContext) SliceIndex() ISliceIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceIndexContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ClauseExpressionParserRULE_primaryExpression)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClauseExpressionParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(ClauseExpressionParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.PrimaryExpression()
		}

	case ClauseExpressionParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(ClauseExpressionParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.GoExpression()
		}
		{
			p.SetState(92)
			p.Match(ClauseExpressionParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ClauseExpressionParserDECIMAL_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Number()
		}

	case ClauseExpressionParserRAW_STRING_LIT, ClauseExpressionParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.String_()
		}

	case ClauseExpressionParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.QualifiedIdentifier()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(97)
				p.FunctionCallArguments()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(98)
				p.SliceIndex()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualifiedIdentifierContext is an interface to support dynamic dispatch.
type IQualifiedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsQualifiedIdentifierContext differentiates from other interfaces.
	IsQualifiedIdentifierContext()
}

type QualifiedIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentifierContext() *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_qualifiedIdentifier
	return p
}

func InitEmptyQualifiedIdentifierContext(p *QualifiedIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_qualifiedIdentifier
}

func (*QualifiedIdentifierContext) IsQualifiedIdentifierContext() {}

func NewQualifiedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_qualifiedIdentifier

	return p
}

func (s *QualifiedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClauseExpressionParserID)
}

func (s *QualifiedIdentifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserID, i)
}

func (s *QualifiedIdentifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ClauseExpressionParserDOT)
}

func (s *QualifiedIdentifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserDOT, i)
}

func (s *QualifiedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterQualifiedIdentifier(s)
	}
}

func (s *QualifiedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitQualifiedIdentifier(s)
	}
}

func (s *QualifiedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitQualifiedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) QualifiedIdentifier() (localctx IQualifiedIdentifierContext) {
	localctx = NewQualifiedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ClauseExpressionParserRULE_qualifiedIdentifier)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(ClauseExpressionParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(104)
				p.Match(ClauseExpressionParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(105)
				p.Match(ClauseExpressionParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallArgumentsContext is an interface to support dynamic dispatch.
type IFunctionCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGoExpression() []IGoExpressionContext
	GoExpression(i int) IGoExpressionContext

	// IsFunctionCallArgumentsContext differentiates from other interfaces.
	IsFunctionCallArgumentsContext()
}

type FunctionCallArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallArgumentsContext() *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_functionCallArguments
	return p
}

func InitEmptyFunctionCallArgumentsContext(p *FunctionCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_functionCallArguments
}

func (*FunctionCallArgumentsContext) IsFunctionCallArgumentsContext() {}

func NewFunctionCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_functionCallArguments

	return p
}

func (s *FunctionCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallArgumentsContext) AllGoExpression() []IGoExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGoExpressionContext); ok {
			len++
		}
	}

	tst := make([]IGoExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGoExpressionContext); ok {
			tst[i] = t.(IGoExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallArgumentsContext) GoExpression(i int) IGoExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoExpressionContext)
}

func (s *FunctionCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterFunctionCallArguments(s)
	}
}

func (s *FunctionCallArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitFunctionCallArguments(s)
	}
}

func (s *FunctionCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitFunctionCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) FunctionCallArguments() (localctx IFunctionCallArgumentsContext) {
	localctx = NewFunctionCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ClauseExpressionParserRULE_functionCallArguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ClauseExpressionParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4026531852) != 0 {
		{
			p.SetState(112)
			p.GoExpression()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ClauseExpressionParserT__17 {
			{
				p.SetState(113)
				p.Match(ClauseExpressionParserT__17)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.GoExpression()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(122)
		p.Match(ClauseExpressionParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceIndexContext is an interface to support dynamic dispatch.
type ISliceIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GoExpression() IGoExpressionContext

	// IsSliceIndexContext differentiates from other interfaces.
	IsSliceIndexContext()
}

type SliceIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceIndexContext() *SliceIndexContext {
	var p = new(SliceIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_sliceIndex
	return p
}

func InitEmptySliceIndexContext(p *SliceIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_sliceIndex
}

func (*SliceIndexContext) IsSliceIndexContext() {}

func NewSliceIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceIndexContext {
	var p = new(SliceIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_sliceIndex

	return p
}

func (s *SliceIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceIndexContext) GoExpression() IGoExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoExpressionContext)
}

func (s *SliceIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterSliceIndex(s)
	}
}

func (s *SliceIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitSliceIndex(s)
	}
}

func (s *SliceIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitSliceIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) SliceIndex() (localctx ISliceIndexContext) {
	localctx = NewSliceIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ClauseExpressionParserRULE_sliceIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(ClauseExpressionParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.GoExpression()
	}
	{
		p.SetState(126)
		p.Match(ClauseExpressionParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_LIT() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserDECIMAL_LIT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ClauseExpressionParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(ClauseExpressionParserDECIMAL_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RAW_STRING_LIT() antlr.TerminalNode
	INTERPRETED_STRING_LIT() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClauseExpressionParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClauseExpressionParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserRAW_STRING_LIT, 0)
}

func (s *StringContext) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ClauseExpressionParserINTERPRETED_STRING_LIT, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClauseExpressionListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClauseExpressionVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClauseExpressionParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ClauseExpressionParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ClauseExpressionParserRAW_STRING_LIT || _la == ClauseExpressionParserINTERPRETED_STRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ClauseExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ClauseExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ClauseExpressionContext)
		}
		return p.ClauseExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ClauseExpressionParser) ClauseExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

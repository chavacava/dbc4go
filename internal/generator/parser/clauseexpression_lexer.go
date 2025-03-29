// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ClauseExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ClauseExpressionLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clauseexpressionlexerLexerInit() {
	staticData := &ClauseExpressionLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'('", "')'", "'=='", "'!='", "'||'", "'&&'", "'>'", "'<'",
		"'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "','", "'['",
		"']'", "'@forall'", "'@exists'", "'==>'", "'@in'", "'@indexof'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "FORALL", "EXISTS", "IMPLIES", "IN", "INDEXOF", "DOT",
		"WHITESPACE", "ID", "DECIMAL_LIT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "FORALL", "EXISTS", "IMPLIES", "IN", "INDEXOF",
		"DOT", "WHITESPACE", "ID", "LETTER", "DECIMAL_LIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 172, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 4, 26, 144, 8, 26, 11, 26, 12, 26, 145, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 5, 27, 153, 8, 27, 10, 27, 12, 27, 156, 9, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 163, 8, 29, 1, 29, 5, 29, 166,
		8, 29, 10, 29, 12, 29, 169, 9, 29, 3, 29, 171, 8, 29, 0, 0, 30, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 0, 59, 29,
		1, 0, 4, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95,
		97, 122, 1, 0, 49, 57, 176, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61,
		1, 0, 0, 0, 3, 63, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9,
		70, 1, 0, 0, 0, 11, 73, 1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15, 79, 1, 0, 0,
		0, 17, 81, 1, 0, 0, 0, 19, 83, 1, 0, 0, 0, 21, 86, 1, 0, 0, 0, 23, 89,
		1, 0, 0, 0, 25, 91, 1, 0, 0, 0, 27, 93, 1, 0, 0, 0, 29, 95, 1, 0, 0, 0,
		31, 97, 1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 101, 1, 0, 0, 0, 37, 103, 1,
		0, 0, 0, 39, 105, 1, 0, 0, 0, 41, 107, 1, 0, 0, 0, 43, 115, 1, 0, 0, 0,
		45, 123, 1, 0, 0, 0, 47, 127, 1, 0, 0, 0, 49, 131, 1, 0, 0, 0, 51, 140,
		1, 0, 0, 0, 53, 143, 1, 0, 0, 0, 55, 149, 1, 0, 0, 0, 57, 157, 1, 0, 0,
		0, 59, 170, 1, 0, 0, 0, 61, 62, 5, 58, 0, 0, 62, 2, 1, 0, 0, 0, 63, 64,
		5, 40, 0, 0, 64, 4, 1, 0, 0, 0, 65, 66, 5, 41, 0, 0, 66, 6, 1, 0, 0, 0,
		67, 68, 5, 61, 0, 0, 68, 69, 5, 61, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71, 5,
		33, 0, 0, 71, 72, 5, 61, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 5, 124, 0, 0,
		74, 75, 5, 124, 0, 0, 75, 12, 1, 0, 0, 0, 76, 77, 5, 38, 0, 0, 77, 78,
		5, 38, 0, 0, 78, 14, 1, 0, 0, 0, 79, 80, 5, 62, 0, 0, 80, 16, 1, 0, 0,
		0, 81, 82, 5, 60, 0, 0, 82, 18, 1, 0, 0, 0, 83, 84, 5, 62, 0, 0, 84, 85,
		5, 61, 0, 0, 85, 20, 1, 0, 0, 0, 86, 87, 5, 60, 0, 0, 87, 88, 5, 61, 0,
		0, 88, 22, 1, 0, 0, 0, 89, 90, 5, 43, 0, 0, 90, 24, 1, 0, 0, 0, 91, 92,
		5, 45, 0, 0, 92, 26, 1, 0, 0, 0, 93, 94, 5, 42, 0, 0, 94, 28, 1, 0, 0,
		0, 95, 96, 5, 47, 0, 0, 96, 30, 1, 0, 0, 0, 97, 98, 5, 37, 0, 0, 98, 32,
		1, 0, 0, 0, 99, 100, 5, 33, 0, 0, 100, 34, 1, 0, 0, 0, 101, 102, 5, 44,
		0, 0, 102, 36, 1, 0, 0, 0, 103, 104, 5, 91, 0, 0, 104, 38, 1, 0, 0, 0,
		105, 106, 5, 93, 0, 0, 106, 40, 1, 0, 0, 0, 107, 108, 5, 64, 0, 0, 108,
		109, 5, 102, 0, 0, 109, 110, 5, 111, 0, 0, 110, 111, 5, 114, 0, 0, 111,
		112, 5, 97, 0, 0, 112, 113, 5, 108, 0, 0, 113, 114, 5, 108, 0, 0, 114,
		42, 1, 0, 0, 0, 115, 116, 5, 64, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118,
		5, 120, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5, 115, 0, 0, 120, 121,
		5, 116, 0, 0, 121, 122, 5, 115, 0, 0, 122, 44, 1, 0, 0, 0, 123, 124, 5,
		61, 0, 0, 124, 125, 5, 61, 0, 0, 125, 126, 5, 62, 0, 0, 126, 46, 1, 0,
		0, 0, 127, 128, 5, 64, 0, 0, 128, 129, 5, 105, 0, 0, 129, 130, 5, 110,
		0, 0, 130, 48, 1, 0, 0, 0, 131, 132, 5, 64, 0, 0, 132, 133, 5, 105, 0,
		0, 133, 134, 5, 110, 0, 0, 134, 135, 5, 100, 0, 0, 135, 136, 5, 101, 0,
		0, 136, 137, 5, 120, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139, 5, 102, 0,
		0, 139, 50, 1, 0, 0, 0, 140, 141, 5, 46, 0, 0, 141, 52, 1, 0, 0, 0, 142,
		144, 7, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 6, 26,
		0, 0, 148, 54, 1, 0, 0, 0, 149, 154, 3, 57, 28, 0, 150, 153, 3, 57, 28,
		0, 151, 153, 7, 1, 0, 0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153,
		156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 56, 1,
		0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 7, 2, 0, 0, 158, 58, 1, 0, 0,
		0, 159, 171, 5, 48, 0, 0, 160, 167, 7, 3, 0, 0, 161, 163, 5, 95, 0, 0,
		162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		166, 7, 1, 0, 0, 165, 162, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0,
		0, 0, 170, 159, 1, 0, 0, 0, 170, 160, 1, 0, 0, 0, 171, 60, 1, 0, 0, 0,
		7, 0, 145, 152, 154, 162, 167, 170, 1, 6, 0, 0,
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

// ClauseExpressionLexerInit initializes any static state used to implement ClauseExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewClauseExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClauseExpressionLexerInit() {
	staticData := &ClauseExpressionLexerLexerStaticData
	staticData.once.Do(clauseexpressionlexerLexerInit)
}

// NewClauseExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewClauseExpressionLexer(input antlr.CharStream) *ClauseExpressionLexer {
	ClauseExpressionLexerInit()
	l := new(ClauseExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ClauseExpressionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ClauseExpression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ClauseExpressionLexer tokens.
const (
	ClauseExpressionLexerT__0        = 1
	ClauseExpressionLexerT__1        = 2
	ClauseExpressionLexerT__2        = 3
	ClauseExpressionLexerT__3        = 4
	ClauseExpressionLexerT__4        = 5
	ClauseExpressionLexerT__5        = 6
	ClauseExpressionLexerT__6        = 7
	ClauseExpressionLexerT__7        = 8
	ClauseExpressionLexerT__8        = 9
	ClauseExpressionLexerT__9        = 10
	ClauseExpressionLexerT__10       = 11
	ClauseExpressionLexerT__11       = 12
	ClauseExpressionLexerT__12       = 13
	ClauseExpressionLexerT__13       = 14
	ClauseExpressionLexerT__14       = 15
	ClauseExpressionLexerT__15       = 16
	ClauseExpressionLexerT__16       = 17
	ClauseExpressionLexerT__17       = 18
	ClauseExpressionLexerT__18       = 19
	ClauseExpressionLexerT__19       = 20
	ClauseExpressionLexerFORALL      = 21
	ClauseExpressionLexerEXISTS      = 22
	ClauseExpressionLexerIMPLIES     = 23
	ClauseExpressionLexerIN          = 24
	ClauseExpressionLexerINDEXOF     = 25
	ClauseExpressionLexerDOT         = 26
	ClauseExpressionLexerWHITESPACE  = 27
	ClauseExpressionLexerID          = 28
	ClauseExpressionLexerDECIMAL_LIT = 29
)

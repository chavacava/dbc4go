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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "FORALL", "EXISTS", "IMPLIES", "IN", "INDEXOF",
		"DOT", "WHITESPACE", "ID", "LETTER", "DECIMAL_LIT", "RAW_STRING_LIT",
		"INTERPRETED_STRING_LIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 194, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 4, 26, 148, 8, 26,
		11, 26, 12, 26, 149, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 157, 8,
		27, 10, 27, 12, 27, 160, 9, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29,
		167, 8, 29, 1, 29, 5, 29, 170, 8, 29, 10, 29, 12, 29, 173, 9, 29, 3, 29,
		175, 8, 29, 1, 30, 1, 30, 5, 30, 179, 8, 30, 10, 30, 12, 30, 182, 9, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 188, 8, 31, 10, 31, 12, 31, 191, 9,
		31, 1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 0, 59, 29, 61, 30, 63, 31, 1, 0, 6, 3, 0, 9, 10,
		13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 49,
		57, 1, 0, 96, 96, 1, 0, 34, 34, 200, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0,
		0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 67, 1, 0, 0,
		0, 5, 69, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 73, 1, 0, 0, 0, 11, 76, 1,
		0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19,
		87, 1, 0, 0, 0, 21, 89, 1, 0, 0, 0, 23, 92, 1, 0, 0, 0, 25, 95, 1, 0, 0,
		0, 27, 97, 1, 0, 0, 0, 29, 99, 1, 0, 0, 0, 31, 101, 1, 0, 0, 0, 33, 103,
		1, 0, 0, 0, 35, 105, 1, 0, 0, 0, 37, 107, 1, 0, 0, 0, 39, 109, 1, 0, 0,
		0, 41, 111, 1, 0, 0, 0, 43, 119, 1, 0, 0, 0, 45, 127, 1, 0, 0, 0, 47, 131,
		1, 0, 0, 0, 49, 135, 1, 0, 0, 0, 51, 144, 1, 0, 0, 0, 53, 147, 1, 0, 0,
		0, 55, 153, 1, 0, 0, 0, 57, 161, 1, 0, 0, 0, 59, 174, 1, 0, 0, 0, 61, 176,
		1, 0, 0, 0, 63, 185, 1, 0, 0, 0, 65, 66, 5, 58, 0, 0, 66, 2, 1, 0, 0, 0,
		67, 68, 5, 33, 0, 0, 68, 4, 1, 0, 0, 0, 69, 70, 5, 40, 0, 0, 70, 6, 1,
		0, 0, 0, 71, 72, 5, 41, 0, 0, 72, 8, 1, 0, 0, 0, 73, 74, 5, 61, 0, 0, 74,
		75, 5, 61, 0, 0, 75, 10, 1, 0, 0, 0, 76, 77, 5, 33, 0, 0, 77, 78, 5, 61,
		0, 0, 78, 12, 1, 0, 0, 0, 79, 80, 5, 124, 0, 0, 80, 81, 5, 124, 0, 0, 81,
		14, 1, 0, 0, 0, 82, 83, 5, 38, 0, 0, 83, 84, 5, 38, 0, 0, 84, 16, 1, 0,
		0, 0, 85, 86, 5, 62, 0, 0, 86, 18, 1, 0, 0, 0, 87, 88, 5, 60, 0, 0, 88,
		20, 1, 0, 0, 0, 89, 90, 5, 62, 0, 0, 90, 91, 5, 61, 0, 0, 91, 22, 1, 0,
		0, 0, 92, 93, 5, 60, 0, 0, 93, 94, 5, 61, 0, 0, 94, 24, 1, 0, 0, 0, 95,
		96, 5, 43, 0, 0, 96, 26, 1, 0, 0, 0, 97, 98, 5, 45, 0, 0, 98, 28, 1, 0,
		0, 0, 99, 100, 5, 42, 0, 0, 100, 30, 1, 0, 0, 0, 101, 102, 5, 47, 0, 0,
		102, 32, 1, 0, 0, 0, 103, 104, 5, 37, 0, 0, 104, 34, 1, 0, 0, 0, 105, 106,
		5, 44, 0, 0, 106, 36, 1, 0, 0, 0, 107, 108, 5, 91, 0, 0, 108, 38, 1, 0,
		0, 0, 109, 110, 5, 93, 0, 0, 110, 40, 1, 0, 0, 0, 111, 112, 5, 64, 0, 0,
		112, 113, 5, 102, 0, 0, 113, 114, 5, 111, 0, 0, 114, 115, 5, 114, 0, 0,
		115, 116, 5, 97, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5, 108, 0, 0,
		118, 42, 1, 0, 0, 0, 119, 120, 5, 64, 0, 0, 120, 121, 5, 101, 0, 0, 121,
		122, 5, 120, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 115, 0, 0, 124,
		125, 5, 116, 0, 0, 125, 126, 5, 115, 0, 0, 126, 44, 1, 0, 0, 0, 127, 128,
		5, 61, 0, 0, 128, 129, 5, 61, 0, 0, 129, 130, 5, 62, 0, 0, 130, 46, 1,
		0, 0, 0, 131, 132, 5, 64, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 110,
		0, 0, 134, 48, 1, 0, 0, 0, 135, 136, 5, 64, 0, 0, 136, 137, 5, 105, 0,
		0, 137, 138, 5, 110, 0, 0, 138, 139, 5, 100, 0, 0, 139, 140, 5, 101, 0,
		0, 140, 141, 5, 120, 0, 0, 141, 142, 5, 111, 0, 0, 142, 143, 5, 102, 0,
		0, 143, 50, 1, 0, 0, 0, 144, 145, 5, 46, 0, 0, 145, 52, 1, 0, 0, 0, 146,
		148, 7, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 6, 26,
		0, 0, 152, 54, 1, 0, 0, 0, 153, 158, 3, 57, 28, 0, 154, 157, 3, 57, 28,
		0, 155, 157, 7, 1, 0, 0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157,
		160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 56, 1,
		0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 7, 2, 0, 0, 162, 58, 1, 0, 0,
		0, 163, 175, 5, 48, 0, 0, 164, 171, 7, 3, 0, 0, 165, 167, 5, 95, 0, 0,
		166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		170, 7, 1, 0, 0, 169, 166, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0,
		0, 0, 174, 163, 1, 0, 0, 0, 174, 164, 1, 0, 0, 0, 175, 60, 1, 0, 0, 0,
		176, 180, 5, 96, 0, 0, 177, 179, 8, 4, 0, 0, 178, 177, 1, 0, 0, 0, 179,
		182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 96, 0, 0, 184, 62, 1, 0,
		0, 0, 185, 189, 5, 34, 0, 0, 186, 188, 8, 5, 0, 0, 187, 186, 1, 0, 0, 0,
		188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		192, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193, 5, 34, 0, 0, 193, 64,
		1, 0, 0, 0, 9, 0, 149, 156, 158, 166, 171, 174, 180, 189, 1, 6, 0, 0,
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
	ClauseExpressionLexerT__0                   = 1
	ClauseExpressionLexerT__1                   = 2
	ClauseExpressionLexerT__2                   = 3
	ClauseExpressionLexerT__3                   = 4
	ClauseExpressionLexerT__4                   = 5
	ClauseExpressionLexerT__5                   = 6
	ClauseExpressionLexerT__6                   = 7
	ClauseExpressionLexerT__7                   = 8
	ClauseExpressionLexerT__8                   = 9
	ClauseExpressionLexerT__9                   = 10
	ClauseExpressionLexerT__10                  = 11
	ClauseExpressionLexerT__11                  = 12
	ClauseExpressionLexerT__12                  = 13
	ClauseExpressionLexerT__13                  = 14
	ClauseExpressionLexerT__14                  = 15
	ClauseExpressionLexerT__15                  = 16
	ClauseExpressionLexerT__16                  = 17
	ClauseExpressionLexerT__17                  = 18
	ClauseExpressionLexerT__18                  = 19
	ClauseExpressionLexerT__19                  = 20
	ClauseExpressionLexerFORALL                 = 21
	ClauseExpressionLexerEXISTS                 = 22
	ClauseExpressionLexerIMPLIES                = 23
	ClauseExpressionLexerIN                     = 24
	ClauseExpressionLexerINDEXOF                = 25
	ClauseExpressionLexerDOT                    = 26
	ClauseExpressionLexerWHITESPACE             = 27
	ClauseExpressionLexerID                     = 28
	ClauseExpressionLexerDECIMAL_LIT            = 29
	ClauseExpressionLexerRAW_STRING_LIT         = 30
	ClauseExpressionLexerINTERPRETED_STRING_LIT = 31
)

// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 155,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 63, 10, 8, 13, 8, 14, 8, 64, 3, 9, 5,
	9, 68, 10, 9, 3, 9, 3, 9, 7, 9, 72, 10, 9, 12, 9, 14, 9, 75, 11, 9, 3,
	10, 6, 10, 78, 10, 10, 13, 10, 14, 10, 79, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 103, 10, 17, 12, 17,
	14, 17, 106, 11, 17, 3, 17, 3, 17, 3, 17, 7, 17, 111, 10, 17, 12, 17, 14,
	17, 114, 11, 17, 5, 17, 116, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23,
	143, 10, 23, 12, 23, 14, 23, 146, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 5, 24, 154, 10, 24, 3, 144, 2, 25, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 2,
	3, 2, 6, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 5, 2, 67, 92, 97,
	97, 99, 124, 7, 2, 48, 48, 50, 59, 67, 92, 97, 97, 99, 124, 2, 163, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 51,
	3, 2, 2, 2, 7, 53, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13,
	59, 3, 2, 2, 2, 15, 62, 3, 2, 2, 2, 17, 67, 3, 2, 2, 2, 19, 77, 3, 2, 2,
	2, 21, 83, 3, 2, 2, 2, 23, 85, 3, 2, 2, 2, 25, 88, 3, 2, 2, 2, 27, 90,
	3, 2, 2, 2, 29, 93, 3, 2, 2, 2, 31, 96, 3, 2, 2, 2, 33, 115, 3, 2, 2, 2,
	35, 117, 3, 2, 2, 2, 37, 121, 3, 2, 2, 2, 39, 124, 3, 2, 2, 2, 41, 128,
	3, 2, 2, 2, 43, 131, 3, 2, 2, 2, 45, 139, 3, 2, 2, 2, 47, 153, 3, 2, 2,
	2, 49, 50, 7, 42, 2, 2, 50, 4, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 6,
	3, 2, 2, 2, 53, 54, 7, 44, 2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7, 49, 2, 2,
	56, 10, 3, 2, 2, 2, 57, 58, 7, 45, 2, 2, 58, 12, 3, 2, 2, 2, 59, 60, 7,
	47, 2, 2, 60, 14, 3, 2, 2, 2, 61, 63, 9, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63,
	64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 16, 3, 2, 2,
	2, 66, 68, 5, 15, 8, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69,
	3, 2, 2, 2, 69, 73, 7, 48, 2, 2, 70, 72, 5, 15, 8, 2, 71, 70, 3, 2, 2,
	2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 18,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 9, 3, 2, 2, 77, 76, 3, 2, 2, 2,
	78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3,
	2, 2, 2, 81, 82, 8, 10, 2, 2, 82, 20, 3, 2, 2, 2, 83, 84, 7, 64, 2, 2,
	84, 22, 3, 2, 2, 2, 85, 86, 7, 64, 2, 2, 86, 87, 7, 63, 2, 2, 87, 24, 3,
	2, 2, 2, 88, 89, 7, 62, 2, 2, 89, 26, 3, 2, 2, 2, 90, 91, 7, 62, 2, 2,
	91, 92, 7, 63, 2, 2, 92, 28, 3, 2, 2, 2, 93, 94, 7, 63, 2, 2, 94, 95, 7,
	63, 2, 2, 95, 30, 3, 2, 2, 2, 96, 97, 7, 35, 2, 2, 97, 98, 7, 63, 2, 2,
	98, 32, 3, 2, 2, 2, 99, 100, 7, 38, 2, 2, 100, 104, 9, 4, 2, 2, 101, 103,
	9, 5, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 116, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	107, 108, 7, 66, 2, 2, 108, 112, 9, 4, 2, 2, 109, 111, 9, 5, 2, 2, 110,
	109, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113,
	3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 99, 3, 2,
	2, 2, 115, 107, 3, 2, 2, 2, 116, 34, 3, 2, 2, 2, 117, 118, 7, 99, 2, 2,
	118, 119, 7, 112, 2, 2, 119, 120, 7, 102, 2, 2, 120, 36, 3, 2, 2, 2, 121,
	122, 7, 113, 2, 2, 122, 123, 7, 116, 2, 2, 123, 38, 3, 2, 2, 2, 124, 125,
	7, 112, 2, 2, 125, 126, 7, 113, 2, 2, 126, 127, 7, 118, 2, 2, 127, 40,
	3, 2, 2, 2, 128, 129, 7, 107, 2, 2, 129, 130, 7, 112, 2, 2, 130, 42, 3,
	2, 2, 2, 131, 132, 7, 41, 2, 2, 132, 133, 5, 15, 8, 2, 133, 134, 5, 13,
	7, 2, 134, 135, 5, 15, 8, 2, 135, 136, 5, 13, 7, 2, 136, 137, 5, 15, 8,
	2, 137, 138, 7, 41, 2, 2, 138, 44, 3, 2, 2, 2, 139, 144, 7, 41, 2, 2, 140,
	143, 5, 47, 24, 2, 141, 143, 11, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 141,
	3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 145, 147, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 7, 41, 2, 2,
	148, 46, 3, 2, 2, 2, 149, 150, 7, 94, 2, 2, 150, 154, 7, 36, 2, 2, 151,
	152, 7, 94, 2, 2, 152, 154, 7, 94, 2, 2, 153, 149, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 154, 48, 3, 2, 2, 2, 13, 2, 64, 67, 73, 79, 104, 112, 115,
	142, 144, 153, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "", "", "", "'>'", "'>='",
	"'<'", "'<='", "'=='", "'!='", "", "'and'", "'or'", "'not'", "'in'",
}

var lexerSymbolicNames = []string{
	"", "", "", "MUL", "DIV", "ADD", "HYPHEN", "DIGIT", "FLOAT", "WHITESPACE",
	"GT", "GE", "LT", "LE", "EQ", "NE", "VARIABLE", "AND", "OR", "NOT", "IN",
	"DATE", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "MUL", "DIV", "ADD", "HYPHEN", "DIGIT", "FLOAT", "WHITESPACE",
	"GT", "GE", "LT", "LE", "EQ", "NE", "VARIABLE", "AND", "OR", "NOT", "IN",
	"DATE", "STRING", "ESC",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewExprLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ExprLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExprLexer(input antlr.CharStream) *ExprLexer {
	l := new(ExprLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0       = 1
	ExprLexerT__1       = 2
	ExprLexerMUL        = 3
	ExprLexerDIV        = 4
	ExprLexerADD        = 5
	ExprLexerHYPHEN     = 6
	ExprLexerDIGIT      = 7
	ExprLexerFLOAT      = 8
	ExprLexerWHITESPACE = 9
	ExprLexerGT         = 10
	ExprLexerGE         = 11
	ExprLexerLT         = 12
	ExprLexerLE         = 13
	ExprLexerEQ         = 14
	ExprLexerNE         = 15
	ExprLexerVARIABLE   = 16
	ExprLexerAND        = 17
	ExprLexerOR         = 18
	ExprLexerNOT        = 19
	ExprLexerIN         = 20
	ExprLexerDATE       = 21
	ExprLexerSTRING     = 22
)

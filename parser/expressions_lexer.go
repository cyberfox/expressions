// Generated from Expressions.g4 by ANTLR 4.7.

package parser
import (
	"fmt"
	"unicode"
)

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 74, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 44, 10, 9, 12, 9, 14, 9, 47, 
	11, 9, 3, 10, 6, 10, 50, 10, 10, 13, 10, 14, 10, 51, 3, 11, 5, 11, 55, 
	10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 61, 10, 12, 12, 12, 14, 12, 
	64, 11, 12, 3, 12, 3, 12, 3, 13, 6, 13, 69, 10, 13, 13, 13, 14, 13, 70, 
	3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2, 17, 
	9, 19, 10, 21, 11, 23, 12, 25, 13, 3, 2, 7, 7, 2, 38, 38, 50, 59, 67, 92, 
	97, 97, 99, 124, 5, 2, 38, 38, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 12, 
	12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 77, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 
	31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 37, 3, 2, 2, 
	2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 49, 3, 2, 2, 2, 21, 54, 
	3, 2, 2, 2, 23, 58, 3, 2, 2, 2, 25, 68, 3, 2, 2, 2, 27, 28, 7, 60, 2, 2, 
	28, 4, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 43, 
	2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 128, 2, 2, 34, 10, 3, 2, 2, 2, 35, 
	36, 7, 45, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 7, 47, 2, 2, 38, 14, 3, 2, 
	2, 2, 39, 40, 9, 2, 2, 2, 40, 16, 3, 2, 2, 2, 41, 45, 9, 3, 2, 2, 42, 44, 
	5, 15, 8, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 
	45, 46, 3, 2, 2, 2, 46, 18, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 50, 9, 
	4, 2, 2, 49, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 
	52, 3, 2, 2, 2, 52, 20, 3, 2, 2, 2, 53, 55, 7, 15, 2, 2, 54, 53, 3, 2, 
	2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 7, 12, 2, 2, 57, 
	22, 3, 2, 2, 2, 58, 62, 7, 37, 2, 2, 59, 61, 10, 5, 2, 2, 60, 59, 3, 2, 
	2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 
	3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 8, 12, 2, 2, 66, 24, 3, 2, 2, 2, 
	67, 69, 9, 6, 2, 2, 68, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 68, 3, 
	2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 8, 13, 2, 2, 73, 
	26, 3, 2, 2, 2, 8, 2, 45, 51, 54, 62, 70, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "'('", "')'", "'~'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "INV", "ADD", "SUB", "ID", "INT", "NEWLINE", "COMMENT", 
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "INV", "ADD", "SUB", "AlphaNum", "ID", "INT", "NEWLINE", 
	"COMMENT", "WS",
}

type ExpressionsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExpressionsLexer(input antlr.CharStream) *ExpressionsLexer {

	l := new(ExpressionsLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expressions.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionsLexer tokens.
const (
	ExpressionsLexerT__0 = 1
	ExpressionsLexerT__1 = 2
	ExpressionsLexerT__2 = 3
	ExpressionsLexerINV = 4
	ExpressionsLexerADD = 5
	ExpressionsLexerSUB = 6
	ExpressionsLexerID = 7
	ExpressionsLexerINT = 8
	ExpressionsLexerNEWLINE = 9
	ExpressionsLexerCOMMENT = 10
	ExpressionsLexerWS = 11
)


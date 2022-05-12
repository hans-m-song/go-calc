package parse

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hans-m-song/go-calc/pkg/util"
)

type TokenChar = string
type SymbolType = string

const (
	symbolTypeNoop       SymbolType = "noop"
	symbolTypeOpCode     SymbolType = "opcode"
	symbolTypeSyntax     SymbolType = "syntax"
	symbolTypeNumber     SymbolType = "number"
	symbolTypeIdentifier SymbolType = "identifier"

	TOKEN_OP_PLUS     TokenChar = "+"
	TOKEN_OP_MINUS    TokenChar = "-"
	TOKEN_OP_DIVIDE   TokenChar = "/"
	TOKEN_OP_MULTIPLY TokenChar = "*"

	TOKEN_SYN_ASSIGN TokenChar = "="
	TOKEN_SYN_LPAREN TokenChar = "("
	TOKEN_SYN_RPAREN TokenChar = ")"
)

var (
	numberMatcher       = regexp.MustCompile("^[0-9]+(\\.[0-9]+)?$")
	alphanumericMatcher = regexp.MustCompile("[A-Za-z0-9]")
	identifierMatcher   = regexp.MustCompile("^[A-Za-z][A-Za-z0-9]*$")
)

type Symbol struct {
	Type  SymbolType
	Match util.MatchFn
}

func (s Symbol) Equals(target Symbol) bool {
	return s.Type == target.Type
}

var (
	SymbolNoop       = Symbol{Type: symbolTypeNoop, Match: func(s string) bool { return s == " " || s == "\n" }}
	SymbolOpPlus     = Symbol{Type: symbolTypeOpCode, Match: func(s string) bool { return s == "+" }}
	SymbolOpMinus    = Symbol{Type: symbolTypeOpCode, Match: func(s string) bool { return s == "-" }}
	SymbolOpDivide   = Symbol{Type: symbolTypeOpCode, Match: func(s string) bool { return s == "/" }}
	SymbolOpMultiply = Symbol{Type: symbolTypeOpCode, Match: func(s string) bool { return s == "*" }}
	SymbolSynAssign  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == "=" }}
	SymbolSynLparen  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == "(" }}
	SymbolSynRparen  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == ")" }}
	SymbolNumber     = Symbol{Type: symbolTypeNumber, Match: numberMatcher.MatchString}
	SymbolIdentifier = Symbol{Type: symbolTypeIdentifier, Match: alphanumericMatcher.MatchString}

	SymbolTable = []Symbol{
		SymbolOpPlus,
		SymbolOpMinus,
		SymbolOpDivide,
		SymbolOpMultiply,
		SymbolSynAssign,
		SymbolSynLparen,
		SymbolSynRparen,
		SymbolNumber,
		SymbolIdentifier,
	}
)

type Token struct {
	Symbol
	Value *string
}

func (t Token) Equals(target Token) bool {
	return t.Symbol.Equals(target.Symbol) && t.Value == target.Value
}

func (t Token) String() string {
	if t.Symbol.Type == symbolTypeNoop {
		return "noop"
	}

	return fmt.Sprintf("%s('%s')", t.Type, *t.Value)
}

var (
	TokenNoop = Token{Symbol: Symbol{Type: symbolTypeNoop}, Value: nil}
)

func NewToken(value string) *Token {
	for _, symbol := range SymbolTable {
		if symbol.Match(value) {
			return &Token{Symbol: symbol, Value: &value}
		}
	}

	return nil
}

type TokenStack struct {
	Tokens []Token
}

func (ts *TokenStack) Push(token Token) {
	ts.Tokens = append(ts.Tokens, token)
}

func (ts TokenStack) Strings() []string {
	serialised := make([]string, len(ts.Tokens))
	for i, token := range ts.Tokens {
		serialised[i] = token.String()
	}

	return serialised
}

func (ts TokenStack) String() string {
	return fmt.Sprintf("[%s]", strings.Join(ts.Strings(), ", "))
}

func MatchOpCodeToken(input string) bool {
	return input == TOKEN_OP_PLUS ||
		input == TOKEN_OP_MINUS ||
		input == TOKEN_OP_DIVIDE ||
		input == TOKEN_OP_MULTIPLY
}

func MatchSyntaxToken(input string) bool {
	return input == TOKEN_SYN_ASSIGN ||
		input == TOKEN_SYN_LPAREN ||
		input == TOKEN_SYN_RPAREN
}

func MatchWhitespaceTokens(input string) bool {
	return input == " " || input == "\n"
}

func MatchNotWhitespaceTokens(input string) bool {
	return !MatchWhitespaceTokens(string(input))
}

func MatchNumber(input string) bool {
	return numberMatcher.MatchString(input)
}

func MatchAlphanumeric(input string) bool {
	return alphanumericMatcher.MatchString(input)
}

func MatchIdentifier(input string) bool {
	return identifierMatcher.MatchString(input)
}

func MatchTokenWordDelimiter(input string) bool {
	return !MatchWhitespaceTokens(input) && !MatchSyntaxToken(input)
}

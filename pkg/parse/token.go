package parse

import (
	"fmt"
	"regexp"
)

type TokenChar = rune
type TokenType = string

const (
	tokenTypeSymbol TokenType = "symbol"
	tokenTypeNumber TokenType = "number"
	tokenTypeAlpha  TokenType = "alpha"

	TOKEN_OP_PLUS     TokenChar = '+'
	TOKEN_OP_MINUS    TokenChar = '-'
	TOKEN_OP_DIVIDE   TokenChar = '/'
	TOKEN_OP_MULTIPLY TokenChar = '*'
	TOKEN_LPAREN      TokenChar = '('
	TOKEN_RPAREN      TokenChar = ')'
)

var (
	numberMatcher = regexp.MustCompile("[0-9.]")
	letterMatcher = regexp.MustCompile("[A-Za-z]")
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s('%s')", t.Type, t.Value)
}

func serializeTokens(tokens []Token) []string {
	serialised := make([]string, len(tokens))
	for i, token := range tokens {
		serialised[i] = token.String()
	}

	return serialised
}

func matchSymbol(input rune) bool {
	return input == TOKEN_OP_PLUS ||
		input == TOKEN_OP_MINUS ||
		input == TOKEN_OP_DIVIDE ||
		input == TOKEN_OP_MULTIPLY ||
		input == TOKEN_LPAREN ||
		input == TOKEN_RPAREN
}

func matchNoop(input rune) bool {
	return input == ' ' ||
		input == '\n'
}

func matchNumber(input rune) bool {
	return numberMatcher.MatchString(string(input))
}

func matchAlpha(input rune) bool {
	return letterMatcher.MatchString(string(input))
}

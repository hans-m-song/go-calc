package parse

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenChar = string
type TokenType = string

const (
	tokenTypeOpCode     TokenType = "opcode"
	tokenTypeSyntax     TokenType = "syntax"
	tokenTypeNumber     TokenType = "number"
	tokenTypeIdentifier TokenType = "identifier"

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

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s('%s')", t.Type, t.Value)
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

func matchOpCode(input string) bool {
	return input == TOKEN_OP_PLUS ||
		input == TOKEN_OP_MINUS ||
		input == TOKEN_OP_DIVIDE ||
		input == TOKEN_OP_MULTIPLY
}

func matchSyntax(input string) bool {
	return input == TOKEN_SYN_ASSIGN ||
		input == TOKEN_SYN_LPAREN ||
		input == TOKEN_SYN_RPAREN
}

func matchWhitespace(input string) bool {
	return input == " " ||
		input == "\n"
}

func matchNumber(input string) bool {
	return numberMatcher.MatchString(input)
}

func matchAlphanumeric(input string) bool {
	return alphanumericMatcher.MatchString(input)
}

func matchIdentifier(input string) bool {
	return identifierMatcher.MatchString(input)
}

func matchWordDelimiter(input string) bool {
	return !matchWhitespace(input) && !matchSyntax(input)
}

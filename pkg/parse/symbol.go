package parse

import "github.com/hans-m-song/go-calc/pkg/util"

type SymbolType = string

const (
	symbolTypeNoop       SymbolType = "noop"
	symbolTypeOperator   SymbolType = "op"
	symbolTypeSyntax     SymbolType = "syntax"
	symbolTypeNumber     SymbolType = "number"
	symbolTypeIdentifier SymbolType = "identifier"

	TOKEN_PLUS     = "+"
	TOKEN_MINUS    = "-"
	TOKEN_DIVIDE   = "/"
	TOKEN_MULTIPLY = "*"
	TOKEN_ASSIGN   = "="
	TOKEN_LPAREN   = "("
	TOKEN_RPAREN   = ")"
	TOKEN_SPACE    = " "
	TOKEN_NEWLINE  = "\n"
)

type Symbol struct {
	Type  SymbolType
	Match util.MatchFn
}

func (s Symbol) Equals(target Symbol) bool {
	return s.Type == target.Type
}

var (
	SymbolNoop       = Symbol{Type: symbolTypeNoop, Match: func(s string) bool { return s == TOKEN_SPACE || s == TOKEN_NEWLINE }}
	SymbolOpPlus     = Symbol{Type: symbolTypeOperator, Match: func(s string) bool { return s == TOKEN_PLUS }}
	SymbolOpMinus    = Symbol{Type: symbolTypeOperator, Match: func(s string) bool { return s == TOKEN_MINUS }}
	SymbolOpDivide   = Symbol{Type: symbolTypeOperator, Match: func(s string) bool { return s == TOKEN_DIVIDE }}
	SymbolOpMultiply = Symbol{Type: symbolTypeOperator, Match: func(s string) bool { return s == TOKEN_MULTIPLY }}
	SymbolSynAssign  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == TOKEN_ASSIGN }}
	SymbolSynLparen  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == TOKEN_LPAREN }}
	SymbolSynRparen  = Symbol{Type: symbolTypeSyntax, Match: func(s string) bool { return s == TOKEN_RPAREN }}
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

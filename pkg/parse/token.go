package parse

import (
	"fmt"
)

type Token struct {
	Symbol   Symbol
	Position int
	Value    *string
}

func (t Token) Equals(target Token) bool {
	return t.Symbol.Equals(target.Symbol) && t.Value == target.Value
}

func (t Token) String() string {
	if t.Symbol.Type == symbolTypeNoop {
		return "noop"
	}

	return fmt.Sprintf("%s('%s')", t.Symbol.Type, *t.Value)
}

var (
	TokenTerminator = Token{Symbol: SymbolTerminator, Value: nil}
	TokenNoop       = Token{Symbol: SymbolNoop, Value: nil}
)

func NewToken(value string, position int) *Token {
	for _, symbol := range SymbolTable {
		if symbol.Match(value) {
			var tokenValue *string = nil
			if !symbol.Equals(SymbolNoop) {
				tokenValue = &value
			}

			return &Token{Symbol: symbol, Position: position, Value: tokenValue}
		}
	}

	return nil
}

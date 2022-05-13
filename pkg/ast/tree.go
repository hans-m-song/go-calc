package ast

import (
	"fmt"

	"github.com/hans-m-song/go-calc/pkg/parse"
)

func BuildAst(ts *parse.TokenStack) (Node, int, error) {
	var (
		current *parse.Token
		next    *parse.Token
		// root     = &ExpressionNode{}
		// current  Node
	)

	for next = ts.Peek(); next != nil; current = ts.Pop() {
		if !current.Symbol.IsFollower(next.Symbol) {
			return nil, current.Position, fmt.Errorf("syntax error at position %d: unexpected token '%s'", current.Position, next.String())
		}

		switch {
		case current.Symbol.Equals(parse.SymbolOpPlus):
		case current.Symbol.Equals(parse.SymbolOpMinus):
		case current.Symbol.Equals(parse.SymbolOpDivide):
		case current.Symbol.Equals(parse.SymbolOpMultiply):
		case current.Symbol.Equals(parse.SymbolSynAssign):
		case current.Symbol.Equals(parse.SymbolSynLparen):
		case current.Symbol.Equals(parse.SymbolSynRparen):
		case current.Symbol.Equals(parse.SymbolNumber):
		case current.Symbol.Equals(parse.SymbolIdentifier):
		case current.Symbol.Equals(parse.SymbolTerminator):
			// successful parse

		default:
			return nil, current.Position, fmt.Errorf("token of unknown symbol at %d: '%s'", current.Position, current.String())
		}
	}

	return nil, current.Position, nil
}

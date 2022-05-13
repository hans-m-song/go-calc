// grammar
//
//  command ::=
//    | expression
//    | identifier := expression
//  expression ::=
//    | term
//    | expression + term
//    | expression - term
//  term ::=
//    | factor
//    | term / factor
//    | term * factor
//  factor ::=
//    | value
//    | + factor
//    | - factor
//    | ( expression )
//  value ::=
//    | number
//    | identifier
//  identifier ::=
//    | alpha

package ast

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hans-m-song/go-calc/pkg/parse"
)

var (
	_ Node = (*BinaryNode)(nil)
	_ Node = (*UnaryNode)(nil)
	_ Node = (*ConstNode)(nil)
)

// TODO use context for variable lookup? Node.Evaluate(context.Context)
type Node interface {
	Evaluate(context.Context) error
	String() string
}

type Command interface{ isCommand() }
type Expression interface{ isExpression() }
type Term interface{ isTerm() }
type Factor interface{ isFactor() }
type Value interface{ isValue() }

// TODO
// type AssignmentNode struct {
// 	Identifier ValueNode
// 	Value      Node
// }

type BinaryNode struct {
	Left  Node
	Op    parse.Token
	Right Node
}

func (n *BinaryNode) isExpression() {}
func (n *BinaryNode) isTerm()       {}

func (n *BinaryNode) Evaluate(context.Context) error { return nil }

func (n *BinaryNode) String() string {
	return fmt.Sprintf("%s %s %s", n.Left.String(), n.Op.String(), n.Right.String())
}

type UnaryNode struct {
	Op    parse.Token
	Value Node
}

func (n *UnaryNode) isFactor() {}

func (n *UnaryNode) Evaluate(context.Context) error { return nil }

func (n *UnaryNode) String() string {
	return fmt.Sprintf("%s%s", *n.Op.Value, n.Value.String())
}

// TODO
// type IdentifierNode struct {
// 	Value ValueNode
// }

type ConstNode struct {
	Value parse.Token
}

func (n *ConstNode) Evaluate(context.Context) error {
	if n.Value.Symbol.Equals(parse.SymbolNumber) {
		strconv.ParseFloat(*n.Value.Value, 64)
	}

	if n.Value.Symbol.Equals(parse.SymbolIdentifier) {

	}

	return fmt.Errorf("ValueNode at position %d resolved to an unexpected token: %s", n.Value.Position, n.Value.String())
}

func (n *ConstNode) isValue() {}

func (n *ConstNode) String() string {
	return *n.Value.Value
}

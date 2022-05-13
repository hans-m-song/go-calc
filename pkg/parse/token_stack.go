package parse

import (
	"fmt"
	"strings"
)

type TokenStack struct {
	Tokens []Token
}

func (ts *TokenStack) Size() int {
	return len(ts.Tokens)
}

func (ts *TokenStack) Peek() *Token {
	if ts.Size() > 1 {
		return &ts.Tokens[1]
	}

	return nil
}

func (ts *TokenStack) Pop() *Token {
	if len(ts.Tokens) > 0 {
		current := ts.Tokens[0]
		ts.Tokens = ts.Tokens[1:]
		return &current
	}

	return nil
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

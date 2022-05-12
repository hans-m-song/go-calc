package parse

import (
	"bytes"
	"fmt"

	"github.com/hans-m-song/go-calc/pkg/util"
	"github.com/rs/zerolog/log"
)

func consumeWord(start string, buffer *util.Buffer, match util.MatchFn) (*Token, int, error) {
	var word string
	var err error

	if word, err = buffer.ConsumeTo(match); err != nil {
		return nil, 0, err
	}

	return NewToken(start + word), len(start + word), nil
}

// Consumes from a buffer and processes result into tokens
//
// Returns a set of tokens and the position where an error ocurred (if any)
func Tokenize(input *bytes.Buffer) (*TokenStack, int, error) {
	if input == nil {
		return nil, 0, fmt.Errorf("no input to read from")
	}

	var (
		buffer   = util.NewBuffer(input)
		err      error
		position = 0
		result   TokenStack
	)

	for len(input.Bytes()) > 0 {
		var (
			width int
			char  string
			token *Token
		)

		if char, err = buffer.Consume(); err != nil {
			return nil, position, fmt.Errorf("failed to read character at position %d: %s", position, err.Error())
		}

		switch {
		case MatchOpCodeToken(char), MatchSyntaxToken(char):
			if token = NewToken(char); err != nil {
				return nil, position, fmt.Errorf("failed to read number at position %d: %s", position, err.Error())
			}

		case MatchNumber(char):
			if token, width, err = consumeWord(char, buffer, MatchNumber); err != nil || token == nil {
				return nil, position, fmt.Errorf("failed to read number at position %d: %s", position, err.Error())
			}

		case MatchIdentifier(char):
			if token, width, err = consumeWord(char, buffer, MatchIdentifier); err != nil || token == nil {
				return nil, position, fmt.Errorf("failed to read identifier at position %d: %s", position, err.Error())
			}

		case MatchWhitespaceTokens(char):
			width = 1
			token = &TokenNoop

		default:
			return nil, position, fmt.Errorf("unhandled character at position %d: '%s'", position, char)
		}

		if token == nil {
			return nil, position, fmt.Errorf("did not match any token at position %d", position)
		}

		position += width
		// discard
		if !token.Equals(TokenNoop) {
			result.Push(*token)
		}
	}

	log.Debug().Strs("tokens", result.Strings()).Send()

	return &result, 0, nil
}

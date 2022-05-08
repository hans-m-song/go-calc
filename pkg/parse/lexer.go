package parse

import (
	"bytes"
	"fmt"

	"github.com/rs/zerolog/log"
)

func lookAhead(input *bytes.Buffer, match func(rune) bool) (bool, error) {
	var r rune
	var err error

	if r, _, err = input.ReadRune(); err != nil {
		return false, err
	}

	if err = input.UnreadRune(); err != nil {
		return false, err
	}

	return match(r), nil
}

func consumeMatchedRunes(start rune, input *bytes.Buffer, match func(rune) bool) (string, error) {
	var r rune
	var err error

	result := string(start)
	for len(input.Bytes()) > 0 {
		var matched bool
		if matched, err = lookAhead(input, match); err != nil {
			return result, err
		}

		if !matched {
			return result, nil
		}

		if r, _, err = input.ReadRune(); err != nil {
			return "", fmt.Errorf("could not read rune from input: %s", err.Error())
		}

		result += string(r)
	}

	return result, nil
}

func Tokenize(input *bytes.Buffer) ([]Token, error) {
	if input == nil {
		return nil, fmt.Errorf("no input to read from")
	}

	result := []Token{}
	var position int
	var r rune
	var value string
	var err error

	for len(input.Bytes()) > 0 {
		if r, _, err = input.ReadRune(); err != nil {
			return nil, fmt.Errorf("could not read rune from input: %s", err.Error())
		}

		switch {
		case matchSymbol(r):
			result = append(result, Token{Type: tokenTypeSymbol, Value: string(r)})
			position += 1

		case matchNumber(r):
			if value, err = consumeMatchedRunes(r, input, matchNumber); err != nil {
				return nil, fmt.Errorf("could not read numbers: %s", err.Error())
			}

			result = append(result, Token{Type: tokenTypeNumber, Value: value})
			position += len(value)

		case matchAlpha(r):
			if value, err = consumeMatchedRunes(r, input, matchAlpha); err != nil {
				return nil, fmt.Errorf("could not read alpha: %s", err.Error())
			}

			result = append(result, Token{Type: tokenTypeAlpha, Value: value})
			position += len(value)

		// ignored
		case matchNoop(r):
			position += 1

		// unhandled
		default:
			return nil, fmt.Errorf("unhandled rune '%c' at position %d", r, position)
		}
	}

	log.Debug().
		Strs("tokens", serializeTokens(result)).
		Msg("parse successful")

	return result, nil
}

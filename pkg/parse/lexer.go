package parse

import (
	"bytes"
	"fmt"

	"github.com/rs/zerolog/log"
)

func createPointerAt(pad int) string {
	pointer := ""
	// offset for prompt
	for len(pointer) < pad {
		pointer += " "
	}
	pointer += "^"
	return pointer
}

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

func consumeMatchedRunes(input *bytes.Buffer, match func(rune) bool) (string, error) {
	var r rune
	var err error

	result := ""
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

func consumeWord(input *bytes.Buffer) (string, int, error) {
	matchNotWhitespace := func(r rune) bool { return !matchWhitespace(string(r)) && !matchSyntax(string(r)) }
	matchOnlyWhitespace := func(r rune) bool { return matchWhitespace(string(r)) }

	var word string
	var whitespace string
	var err error
	if word, err = consumeMatchedRunes(input, matchNotWhitespace); err != nil {
		return word, len(word), err
	}

	if _, err = consumeMatchedRunes(input, matchOnlyWhitespace); err != nil {
		return word, len(word), err
	}

	return word, len(word + whitespace), nil
}

func Tokenize(input *bytes.Buffer) (*TokenStack, error) {
	if input == nil {
		return nil, fmt.Errorf("no input to read from")
	}

	result := new(TokenStack)
	var position int
	var word string
	var width int
	var err error

	for len(input.Bytes()) > 0 {
		if word, width, err = consumeWord(input); err != nil {
			return nil, fmt.Errorf("failed to read word: %s", err.Error())
		}

		fmt.Println("word", word)

		switch {
		case matchOpCode(word):
			result.Push(Token{Type: tokenTypeOpCode, Value: string(word)})

		case matchSyntax(word):
			result.Push(Token{Type: tokenTypeSyntax, Value: string(word)})

		case matchIdentifier(word):
			result.Push(Token{Type: tokenTypeIdentifier, Value: word})

		case matchNumber(word):
			result.Push(Token{Type: tokenTypeNumber, Value: word})

		// unhandled
		default:
			// offset +2 for prompt
			fmt.Println(createPointerAt(position + 2))
			return nil, fmt.Errorf("unhandled word '%s' at position %d", word, position)
		}

		position += width
	}

	log.Debug().
		Strs("tokens", result.Strings()).
		Send()

	return result, nil
}

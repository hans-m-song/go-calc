package util

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	buffer *bytes.Buffer
}

func NewBuffer(input *bytes.Buffer) *Buffer {
	return &Buffer{input}
}

type MatchFn func(string) bool

func (b *Buffer) PeekRune(match MatchFn) (bool, error) {
	var r rune
	var err error

	if r, _, err = b.buffer.ReadRune(); err != nil {
		return false, err
	}

	if err = b.buffer.UnreadRune(); err != nil {
		return false, err
	}

	return match(string(r)), nil
}

func (b *Buffer) Consume() (string, error) {
	r, _, err := b.buffer.ReadRune()
	return string(r), err
}

func (b *Buffer) ConsumeTo(matchers ...MatchFn) (string, error) {
	if len(matchers) < 1 {
		return "", fmt.Errorf("must provide at least one match condition")
	}

	var r rune
	var err error
	result := ""
	for len(b.buffer.Bytes()) > 0 {
		var matched bool
		for _, matcher := range matchers {
			if matched, err = b.PeekRune(matcher); err != nil || !matched {
				return result, err
			}
		}

		if r, _, err = b.buffer.ReadRune(); err != nil {
			return result, fmt.Errorf("could not read rune from input: %s", err.Error())
		}

		result += string(r)
	}

	return result, nil
}

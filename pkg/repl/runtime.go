package repl

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/hans-m-song/go-calc/pkg/parse"
)

const (
	promptStr = "# "
)

func createPointerAt(pad int) string {
	pointer := ""
	// offset for prompt
	for len(promptStr+pointer) < pad {
		pointer += " "
	}
	pointer += "^"
	return pointer
}

func read(ctx context.Context, reader *bufio.Reader) (*bytes.Buffer, error) {
	var raw bytes.Buffer
	var err error

	for {
		select {
		case <-ctx.Done():
			return nil, nil

		default:
			var r rune
			if r, _, err = reader.ReadRune(); err != nil {
				return nil, err
			}

			if r == '\n' {
				return &raw, nil
			}

			if _, err = raw.WriteRune(r); err != nil {
				return nil, err
			}
		}
	}
}

func Repl(ctx context.Context, input *os.File) error {
	reader := bufio.NewReader(input)
	var raw *bytes.Buffer
	var err error

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			// prompt
			fmt.Print(promptStr)

			// read
			// TODO handle up for cmd history?
			if raw, err = read(ctx, reader); err != nil || raw == nil {
				return err
			}

			// parse
			var tokens *parse.TokenStack
			var pointer int
			if tokens, pointer, err = parse.Tokenize(raw); err != nil || tokens == nil {
				fmt.Println(createPointerAt(pointer))
				fmt.Printf("could not process input: %s\n", err.Error())
				continue
			}

			// TODO interpret

			// TODO print
		}

		// loop
	}
}

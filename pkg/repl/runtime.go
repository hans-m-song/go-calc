package repl

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/hans-m-song/go-calc/pkg/parse"
	"github.com/rs/zerolog/log"
)

func isStopChar(stop []rune, v rune) bool {
	for _, r := range stop {
		if r == v {
			return true
		}
	}
	return false
}

func read(ctx context.Context, reader *bufio.Reader, stop []rune) (*bytes.Buffer, error) {
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

			if isStopChar(stop, r) {
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
	stop := []rune{'\n'}
	var raw *bytes.Buffer
	var err error

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			// prompt
			fmt.Print("# ")

			// read
			// TODO handle up for cmd history?
			if raw, err = read(ctx, reader, stop); err != nil || raw == nil {
				return err
			}

			// parse
			var tokens []parse.Token
			if tokens, err = parse.Tokenize(raw); err != nil || tokens == nil {
				log.Err(err).Msg("could not process input")
				continue // ignore
			}

			// TODO interpret

			// TODO print
		}

		// loop
	}
}

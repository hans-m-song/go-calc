package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/hans-m-song/go-calc/pkg/repl"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:             os.Stdout,
		FormatTimestamp: func(i interface{}) string { return "" },
	})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)

	cleanup := func() {
		signal.Stop(interupt)
		cancel()
		close(interupt)
		os.Exit(0)
	}

	go func() {
		// listen for an interupt or context cancellation
		select {
		case <-ctx.Done():
		case <-interupt:
			cancel()
		}

		cleanup()
	}()

	if err := repl.Repl(ctx, os.Stdin); err != nil {
		fmt.Println(err.Error())
	}
}

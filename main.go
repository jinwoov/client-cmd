package main

import (
	"client/gurl"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := gurl.CreateCommand().Execute(); err != nil {
		switch e := err.(type) {
		case gurl.ReturnCodeError:
			os.Exit(e.Code())
		default:
			os.Exit(1)
		}
	}
}

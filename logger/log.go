package logger

import (
	"github.com/nnqq/scr-parser/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

var Log zerolog.Logger

func init() {
	level, err := zerolog.ParseLevel(config.Env.LogLevel)
	if err != nil {
		panic(err)
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(level)
}

func Must(err error) {
	if err != nil {
		Log.Panic().Err(err).Send()
	}
}

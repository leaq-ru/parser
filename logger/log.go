package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

var Log zerolog.Logger

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func Must(err error) {
	if err != nil {
		Log.Panic().Err(err).Send()
	}
}

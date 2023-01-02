package framework

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = &logger{}

type logger struct {
}

func (*logger) New() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func (*logger) Info(msg string) {
	log.Info().Msg(msg)
}

func (*logger) Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

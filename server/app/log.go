package app

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = newLogger()

type logger struct {
}

func newLogger() *logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &logger{}
}

func (*logger) Info(msg string) {
	log.Info().Msg(msg)
}

func (*logger) Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

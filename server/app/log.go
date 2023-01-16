package app

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

func (*logger) InfoExtra() *zerolog.Event {
	return log.Info()
}

func (*logger) Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

func (*logger) ErrorWothStack(err error, msg string) {
	log.Error().Stack().Err(err).Msg(msg)
}

func (*logger) ErrorExtra(err error) *zerolog.Event {
	return log.Error().Err(err)
}

func (*logger) Panic(err error, msg string) {
	log.Panic().Err(err).Msg(msg)
}

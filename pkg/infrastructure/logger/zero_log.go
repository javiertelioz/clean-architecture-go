package logger

import (
	"fmt"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"runtime"
)

type ZerologLogger struct{}

func NewLogger() services.LoggerService {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &ZerologLogger{}
}

func (z *ZerologLogger) Trace(msg string) {
	_, file, line, _ := runtime.Caller(1)

	log.Trace().
		Str("loc", fmt.Sprintf("%s:%s", file, string(rune(line)))).
		Msg(msg)

}

func (z *ZerologLogger) Debug(msg string) {
	_, file, line, _ := runtime.Caller(1)
	log.Debug().
		Str("loc", fmt.Sprintf("%s:%s", file, string(rune(line)))).
		Msg(msg)
}

func (z *ZerologLogger) Info(msg string) {
	_, file, line, _ := runtime.Caller(1)
	log.Info().
		Str("loc", fmt.Sprintf("%s:%s", file, string(rune(line)))).
		Msg(msg)
}

func (z *ZerologLogger) Warn(msg string) {
	_, file, line, _ := runtime.Caller(1)
	log.Warn().
		Str("loc", fmt.Sprintf("%s:%s", file, string(rune(line)))).
		Msg(msg)
}

func (z *ZerologLogger) Error(msg string) {
	_, file, line, _ := runtime.Caller(1)
	log.Error().
		Stack().
		Str("loc", fmt.Sprintf("%s:%s", file, string(rune(line)))).
		Msg(msg)
}

package logger

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
)

// ZapLogger godoc
type ZapLogger struct {
	logger *zap.Logger
}

// NewLogger godoc
func NewLogger() services.LoggerService {
	config := zap.Config{
		Encoding:         "json", // O "console" si prefieres texto plano
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			MessageKey:    "message",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize zap logger: %v", err))
	}

	return &ZapLogger{logger: logger}
}

func (z *ZapLogger) getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2) // 2 niveles arriba para obtener la llamada original
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (z *ZapLogger) Trace(msg string) {
	z.logger.Debug(msg, zap.String("loc", z.getCallerInfo()))
}

func (z *ZapLogger) Debug(msg string) {
	z.logger.Debug(msg, zap.String("loc", z.getCallerInfo()))
}

func (z *ZapLogger) Info(msg string) {
	z.logger.Info(msg, zap.String("loc", z.getCallerInfo()))
}

func (z *ZapLogger) Warn(msg string) {
	z.logger.Warn(msg, zap.String("loc", z.getCallerInfo()))
}

func (z *ZapLogger) Error(msg string) {
	z.logger.Error(msg, zap.String("loc", z.getCallerInfo()), zap.Stack("stacktrace"))
}

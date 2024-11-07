package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	_logFieldService   = "service"
	_logFieldComponent = "component"
)

var (
	defaultLogger = newDefaultLogger()

	stringLevel2ZeroLoggerLevel = map[string]zerolog.Level{
		"info":  zerolog.InfoLevel,
		"warn":  zerolog.WarnLevel,
		"error": zerolog.ErrorLevel,
		"debug": zerolog.DebugLevel,
	}
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

type ZeroLoggerWrapper struct {
	logger *zerolog.Logger
}

// Ensure ZeroLoggerWrapper implements Logger
var _ Logger = (*ZeroLoggerWrapper)(nil)

func (z *ZeroLoggerWrapper) Debug(msg string, fields ...Field) {
	z.log(z.logger.Debug(), msg, fields...)
}

func (z *ZeroLoggerWrapper) Info(msg string, fields ...Field) {
	z.log(z.logger.Info(), msg, fields...)
}

func (z *ZeroLoggerWrapper) Warn(msg string, fields ...Field) {
	z.log(z.logger.Warn(), msg, fields...)
}

func (z *ZeroLoggerWrapper) Error(msg string, fields ...Field) {
	z.log(z.logger.Error(), msg, fields...)
}

func (z *ZeroLoggerWrapper) log(event *zerolog.Event, msg string, fields ...Field) {
	for _, field := range fields {
		field(event)
	}
	event.Msg(msg)
}

func NewLogger(serviceName string, level string) *ZeroLoggerWrapper {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(stringLevel2ZeroLoggerLevel[level])
	logger := log.With().Str(_logFieldService, serviceName).Logger()

	wrapper := &ZeroLoggerWrapper{
		logger: &logger,
	}

	defaultLogger = wrapper

	return wrapper
}

func newDefaultLogger() *ZeroLoggerWrapper {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.With().Str(_logFieldService, "default").Logger()

	return &ZeroLoggerWrapper{
		logger: &logger,
	}
}

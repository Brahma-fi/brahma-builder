package log

import (
	"fmt"

	"github.com/rs/zerolog"
	tlog "go.temporal.io/sdk/log"
)

type TemporalLogger struct {
	logger *zerolog.Logger
}

// Ensure TemporalLogger implements tlog.Logger
var _ tlog.Logger = (*TemporalLogger)(nil)

func NewTemporalLogger(serviceName string, level string) *TemporalLogger {
	zLogger := NewLogger(serviceName, level)
	return &TemporalLogger{
		logger: zLogger.logger,
	}
}

func NewTemporalLoggerFromExisting(logger Logger) *TemporalLogger {
	if zLogger, ok := logger.(*ZeroLoggerWrapper); ok {
		return &TemporalLogger{
			logger: zLogger.logger,
		}
	}
	// If it's not a ZeroLoggerWrapper, create a new default logger
	return NewTemporalLogger("default", "info")
}

func (t *TemporalLogger) Debug(msg string, keyvals ...interface{}) {
	t.log(t.logger.Debug(), msg, keyvals...)
}

func (t *TemporalLogger) Info(msg string, keyvals ...interface{}) {
	t.log(t.logger.Info(), msg, keyvals...)
}

func (t *TemporalLogger) Warn(msg string, keyvals ...interface{}) {
	t.log(t.logger.Warn(), msg, keyvals...)
}

func (t *TemporalLogger) Error(msg string, keyvals ...interface{}) {
	t.log(t.logger.Error(), msg, keyvals...)
}

func (t *TemporalLogger) log(event *zerolog.Event, msg string, keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			event.Interface(fmt.Sprintf("%v", keyvals[i]), keyvals[i+1])
		}
	}
	event.Msg(msg)
}

func (t *TemporalLogger) With(keyvals ...interface{}) tlog.Logger {
	newLogger := t.logger.With().Logger()
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			newLogger = newLogger.With().Interface(fmt.Sprintf("%v", keyvals[i]), keyvals[i+1]).Logger()
		}
	}
	return &TemporalLogger{
		logger: &newLogger,
	}
}

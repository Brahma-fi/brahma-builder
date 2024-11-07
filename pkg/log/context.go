package log

import (
	"context"
)

type loggerCtxKey struct{}

func SetLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func SetLoggerWithReqID(ctx context.Context, logger Logger, id string) context.Context {
	wrapper, ok := logger.(*ZeroLoggerWrapper)
	if !ok {
		logger.Error("failed to set request id, invalid logger type")
		return ctx
	}

	newLogger := wrapper.logger.With().Logger()

	newCtx := context.WithValue(ctx, loggerCtxKey{}, &ZeroLoggerWrapper{
		logger: &newLogger,
	})

	return newCtx
}

func GetLogger(ctx context.Context) Logger {
	val := ctx.Value(loggerCtxKey{})
	if val == nil {
		return defaultLogger
	}

	logger, ok := val.(Logger)
	if !ok {
		return defaultLogger
	}

	return logger
}

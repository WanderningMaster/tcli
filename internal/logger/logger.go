package logger

import (
	"context"
)

type loggerKey struct{}
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
}

func WithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

func FromContext(ctx context.Context) Logger {
	logger := ctx.Value(loggerKey{})

	return logger.(Logger)
}

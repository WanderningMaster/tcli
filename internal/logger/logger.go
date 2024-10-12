package logger

import (
	"context"
)

type loggerKey struct{}
type Logger interface {
	Info(msg string, args ...any)
	Warn(format string, a ...any)
	Error(format string, a ...any)
	Debug(format string, a ...any)
}

func WithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

func FromContext(ctx context.Context) Logger {
	logger := ctx.Value(loggerKey{})

	return logger.(Logger)
}

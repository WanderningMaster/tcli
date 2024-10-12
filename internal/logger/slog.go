package logger

import (
	"log/slog"
	"os"
)

func NewSlog() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return logger
}

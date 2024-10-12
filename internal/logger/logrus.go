package logger

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

type PlainFormatter struct {
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

func NewLogrus(debug bool) *log.Logger {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{})

	if !debug {
		logger.SetOutput(io.Discard)
	}

	return logger
}

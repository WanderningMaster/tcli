package internal

import (
	"context"
	"log"
	"os"
)

func GetHomeDir(ctx context.Context) string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

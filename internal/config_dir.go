package internal

import (
	"context"
	"log"
	"os"
)

func GetConfigDir(ctx context.Context) string {
	dirname, err := os.UserConfigDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

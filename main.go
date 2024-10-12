package main

import (
	"context"
	"os"

	"github.com/WanderningMaster/tcli/cmd"
	"github.com/WanderningMaster/tcli/config"
	"github.com/WanderningMaster/tcli/internal/encoding"
	"github.com/WanderningMaster/tcli/internal/logger"
)

var debug = true

func main() {
	log := logger.NewLogrus(debug)
	ctx := logger.WithLogger(context.Background(), log)

	tomlParser := encoding.NewTomlParser()

	cfg := config.NewConfig(ctx, tomlParser)
	if cfg == nil {
		os.Exit(1)
	}

	cmd.Execute(ctx, cfg)
}

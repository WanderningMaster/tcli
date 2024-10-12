package main

import (
	"context"
	"os"

	"github.com/WanderningMaster/tcli/cmd"
	"github.com/WanderningMaster/tcli/config"
	"github.com/WanderningMaster/tcli/internal/encoding"
	"github.com/WanderningMaster/tcli/internal/logger"
)

func main() {
	log := logger.NewSlog()
	ctx := logger.WithLogger(context.Background(), log)

	tomlParser := encoding.NewTomlParser()

	cfg := config.NewConfig(ctx, tomlParser)
	if cfg == nil {
		os.Exit(1)
	}

	cmd.Execute(ctx, cfg)
}

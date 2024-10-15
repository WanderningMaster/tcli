package cmd

import (
	"context"
	"os"

	"github.com/WanderningMaster/tcli/config"
	"github.com/WanderningMaster/tcli/internal/encoding"
	"github.com/WanderningMaster/tcli/internal/infrastructure"
	"github.com/WanderningMaster/tcli/internal/logger"
	"github.com/spf13/cobra"
)

const debug = true

var rootCmd = &cobra.Command{
	Use:   "tcli",
	Short: "Task CLI is simple task manager right in your terminal",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log := logger.NewLogrus(debug)
		ctx := logger.WithLogger(context.Background(), log)

		p := encoding.NewTomlParser()

		cfg := config.NewConfig(ctx, p)
		if cfg == nil {
			os.Exit(1)
		}

		bp := encoding.NewBinaryParser()
		storage := infrastructure.NewStorage(ctx, cfg.StoragePath, bp)
		err := storage.LoadTasks(ctx)

		if err != nil {
			return err
		}

		ctx = infrastructure.WithStorage(ctx, storage)

		cmd.SetContext(ctx)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

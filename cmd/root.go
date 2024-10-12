package cmd

import (
	"context"
	"os"

	"github.com/WanderningMaster/tcli/config"
	"github.com/WanderningMaster/tcli/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tcli",
	Short: "Task CLI is simple task manager right in your terminal",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute(ctx context.Context, cfg *config.Config) {
	logger := logger.FromContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

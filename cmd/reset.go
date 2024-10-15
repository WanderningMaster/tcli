package cmd

import (
	"fmt"

	"github.com/WanderningMaster/tcli/internal/infrastructure"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		storage := infrastructure.GetStorage(ctx)

		err := storage.Reset(ctx)
		if err != nil {
			return err
		}

		fmt.Println("Task successfully reseted!")

		return nil
	},
}

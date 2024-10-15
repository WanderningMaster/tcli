package cmd

import (
	"fmt"
	"strconv"

	"github.com/WanderningMaster/tcli/internal/infrastructure"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove task from the list",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		ctx := cmd.Context()
		storage := infrastructure.GetStorage(ctx)

		err = storage.Remove(ctx, id)
		if err != nil {
			return err
		}

		fmt.Println("Task successfully removed!")

		return nil
	},
}

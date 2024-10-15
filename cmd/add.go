package cmd

import (
	"fmt"
	"strings"

	"github.com/WanderningMaster/tcli/internal/infrastructure"
	"github.com/spf13/cobra"
)

var tag string

func init() {
	addCmd.Flags().StringVarP(&tag, "tag", "t", "all", "Specify Tag for the task")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to the list",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := strings.Join(args, " ")

		ctx := cmd.Context()
		storage := infrastructure.GetStorage(ctx)

		err := storage.Add(ctx, tag, content)
		if err != nil {
			return err
		}

		fmt.Println("Task successfully added!")

		return nil
	},
}

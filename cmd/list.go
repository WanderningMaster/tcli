package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/WanderningMaster/tcli/internal/infrastructure"
	"github.com/WanderningMaster/tcli/internal/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().StringVarP(&tag, "tag", "t", "", "Specify Tag for grouping task list")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all current tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		storage := infrastructure.GetStorage(ctx)

		var tasks []*model.Task
		var err error
		if tag != "" {
			tasks, err = storage.TasksByTag(ctx, tag)
		} else {
			tasks, err = storage.Tasks(ctx)
		}

		if errors.Is(err, infrastructure.StorageEmpty) {
			fmt.Println("Your list is empty")
			return nil
		}
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Tag", "Task"})

		for idx, task := range tasks {
			t.AppendRow(table.Row{
				idx + 1,
				task.Tag,
				task.Content,
			})
		}
		t.Render()

		return nil
	},
}

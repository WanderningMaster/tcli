package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TCLI v0.0.1")
	},
}

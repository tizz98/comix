package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the program version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

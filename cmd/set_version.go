package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tizz98/comix/cnc"
)

var setVersionCmd = &cobra.Command{
	Use: "set-version",
	RunE: func(cmd *cobra.Command, args []string) error {
		version, _ := cmd.Flags().GetInt("version")
		service, err := cnc.New()

		if err != nil {
			return err
		}

		return service.SetLatestFileVersion(version)
	},
}

func init() {
	rootCmd.AddCommand(setVersionCmd)

	setVersionCmd.Flags().Int("version", 0, "the version number (e.g. 1, 2, etc.)")
	setVersionCmd.MarkFlagRequired("version")
}

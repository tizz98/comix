package cmd

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tizz98/comix/app"
)

var downloaderCmd = &cobra.Command{
	Use:   "downloader",
	Short: "Run comix downloader asynchronously.",
	Long:  "Runs the comix downloader asynchronously, to stop press ctrl+c, the downloader will clean up itself.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var downloaderType app.DownloaderType

		if val, _ := cmd.Flags().GetString("source"); val == "xkcd" {
			downloaderType = app.DownloaderTypeXkcd
		}

		outputDir, _ := cmd.Flags().GetString("output-dir")

		logrus.Info("Starting downloader...")
		return app.RunDownloader(downloaderType, outputDir, &app.Option{TickDuration: 10 * time.Second})
	},
}

func init() {
	rootCmd.AddCommand(downloaderCmd)

	downloaderCmd.Flags().String("source", "", "the source of comics to download (e.g. xkcd)")
	downloaderCmd.MarkFlagRequired("source")

	downloaderCmd.Flags().String("output-dir", ".", "the directory to save the comic image files")
}

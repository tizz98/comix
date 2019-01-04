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

		downloadSrc, _ := cmd.Flags().GetString("source")
		switch downloadSrc {
		case "xkcd":
			downloaderType = app.DownloaderTypeXkcd
		}

		outputDir, _ := cmd.Flags().GetString("output-dir")

		logrus.Info("Starting downloader...")
		options := &app.Option{TickDuration: app.NewDuration(10 * time.Second)}

		if cncAddr, _ := cmd.Flags().GetString("cnc"); cncAddr != "" {
			options.CnCAddress = cncAddr
			options.ClientId, _ = cmd.Flags().GetString("client-id")
		}

		return app.RunDownloader(downloaderType, outputDir, options)
	},
}

func init() {
	rootCmd.AddCommand(downloaderCmd)

	downloaderCmd.Flags().String("source", "", "the source of comics to download (e.g. xkcd)")
	downloaderCmd.MarkFlagRequired("source")

	downloaderCmd.Flags().String("client-id", "", "the id of this client (e.g. 123, foo-bar)")

	downloaderCmd.Flags().String("output-dir", ".", "the directory to save the comic image files")
	downloaderCmd.Flags().String("cnc", "", "an optional cnc server (e.g. 127.0.0.1:1337)")
}

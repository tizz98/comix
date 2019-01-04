package cmd

import (
	"os"
	"time"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Comix program",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !terminal.IsTerminal(int(os.Stdout.Fd())) {
			logrus.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: time.RFC3339Nano,
			})
		} else {
			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
			})
		}

		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		viper.SetEnvPrefix("comix")
		viper.AutomaticEnv()
	},
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

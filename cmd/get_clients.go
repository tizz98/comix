package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tizz98/comix/cnc"
)

var getClientsCmd = &cobra.Command{
	Use: "get-clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := cnc.New()
		if err != nil {
			return err
		}

		clients, err := service.GetClients()
		if err != nil {
			return err
		}

		for _, c := range clients {
			logrus.WithFields(logrus.Fields{
				"id":          c.Id,
				"ok":          c.Ok,
				"message":     c.Message,
				"last_update": c.LastUpdate,
			}).Info("")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getClientsCmd)
}

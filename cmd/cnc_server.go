package cmd

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/tizz98/comix/cnc"
)

var cncServerCmd = &cobra.Command{
	Use: "cnc",
	RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetInt("port")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return errors.Wrap(err, "unable to start listener")
		}

		srv := grpc.NewServer()
		service, err := cnc.New()

		if err != nil {
			return errors.Wrap(err, "unable to create new service")
		}
		cnc.RegisterCnCServer(srv, service)

		go srv.Serve(lis)
		logrus.Infof("CNC Service listing on 0.0.0.0:%d", port)

		signalChan := make(chan os.Signal, 1)
		done := make(chan struct{})
		signal.Notify(signalChan, os.Interrupt)

		go func() {
			for {
				select {
				case <-signalChan:
					logrus.Info("received interrupt, shutting down...")
					srv.GracefulStop()
					close(done)
				}
			}
		}()

		<-done
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cncServerCmd)

	cncServerCmd.Flags().Int("port", 1337, "the port to run the grpc server on")
}

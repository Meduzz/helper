package starters

import (
	"os"
	"os/signal"

	"github.com/Meduzz/helper/nuts"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func Nats(setup func(*nats.Conn)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start the nats handler",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		setup(conn)

		// block
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		conn.Close()

		return nil
	}

	return cmd
}

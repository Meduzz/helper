package starters

import (
	"github.com/Meduzz/helper/block"
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

		defer conn.Close()

		return block.Block(func() error {
			return conn.Drain()
		})
	}

	return cmd
}

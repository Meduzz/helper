package starters

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	wendyrpc "github.com/Meduzz/wendy-rpc"
	"github.com/spf13/cobra"
)

func WendyModules(prefix string, modules ...*wendy.Module) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a few wendy module on top of a nats conn",
	}

	cmd.Flags().String("q", "", "set queue group or leave empty")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		queue, err := cmd.Flags().GetString("q")

		if err != nil {
			return err
		}

		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		err = wendyrpc.ServeModules(conn, queue, prefix, modules...)

		if err != nil {
			return err
		}

		defer conn.Close()

		return block.Block(func() error {
			return conn.Drain()
		})
	}

	return cmd
}

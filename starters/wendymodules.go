package starters

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	wendyrpc "github.com/Meduzz/wendy-rpc"
	"github.com/spf13/cobra"
)

func WendyModules(modules ...*wendy.Module) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a few wendy module on top of a nats conn",
	}

	queue := cmd.Flags().StringP("queue", "q", "", "set queue group or leave empty")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		err = wendyrpc.ServeModules(conn, *queue, modules...)

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

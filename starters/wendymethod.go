package starters

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	wendyrpc "github.com/Meduzz/wendy-rpc"
	"github.com/spf13/cobra"
)

func WendyMethod(module, method string, handler wendy.Handler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a wendy method",
	}

	queue := cmd.Flags().StringP("queue", "q", "", "set queue group for nats topic")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		err = wendyrpc.ServeMethod(conn, *queue, module, method, handler)

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

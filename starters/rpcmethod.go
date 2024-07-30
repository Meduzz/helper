package starters

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/encoding"
	"github.com/spf13/cobra"
)

func RpcMethod(handler rpc.RpcHandler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a rpc method",
	}

	topic := cmd.Flags().StringP("topic", "t", "", "topic to bind this rpc method")
	queue := cmd.Flags().StringP("queue", "q", "", "queue to bind this rpc method or leave empty")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		srv := rpc.NewRpc(conn, encoding.Json())

		srv.HandleRPC(*topic, *queue, handler)

		srv.Run()

		return nil
	}

	return cmd
}

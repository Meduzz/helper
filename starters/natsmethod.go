package starters

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func NatsMethod(handler nats.MsgHandler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start the nats handler",
	}

	topic := cmd.Flags().StringP("topic", "t", "", "the topic to bind to")
	queue := cmd.Flags().StringP("queue", "q", "", "the queue group to use")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		if *queue != "" {
			conn.QueueSubscribe(*topic, *queue, handler)
		} else {
			conn.Subscribe(*topic, handler)
		}

		defer conn.Close()

		return block.Block(func() error {
			return conn.Drain()
		})
	}

	return cmd
}

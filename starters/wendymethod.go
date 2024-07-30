package starters

import (
	"fmt"

	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	wendyrpc "github.com/Meduzz/wendy-rpc"
	"github.com/spf13/cobra"
)

func WendyMethod(app, module, method string, handler wendy.Handler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a wendy method",
	}

	cmd.Flags().String("q", "", "set queue group for nats topic")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		queue, err := cmd.Flags().GetString("q")

		if err != nil {
			return err
		}

		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		topic := fmt.Sprintf("%s.%s.%s", app, module, method)

		if app == "" {
			topic = fmt.Sprintf("%s.%s", module, method)
		}

		err = wendyrpc.ServeMethod(conn, queue, topic, handler)

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

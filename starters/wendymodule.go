package starters

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func WendyModule(module *wendy.Module) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a wendy module on top of a nats conn",
	}

	queue := cmd.Flags().StringP("queue", "q", "", "set queue group or leave empty")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		if *queue != "" {
			conn.QueueSubscribe(fmt.Sprintf("%s.*", module.Name()), *queue, wrapModule(module))
		} else {
			conn.Subscribe(fmt.Sprintf("%s.*", module.Name()), wrapModule(module))
		}

		defer conn.Close()

		return block.Block(func() error {
			return conn.Drain()
		})
	}

	return cmd
}

func wrapModule(module *wendy.Module) nats.MsgHandler {
	return func(msg *nats.Msg) {
		req := &wendy.Request{}
		err := json.Unmarshal(msg.Data, req)

		if err != nil {
			log.Printf("[message handler] data could not be parsed to wendy request: %s", msg.Subject)
			return
		}

		h := module.Method(req.Method)

		if h == nil {
			log.Printf("[message handler] no handler found for %s", req.Method)
			return
		}

		res := h(req)

		bs, err := json.Marshal(res)

		if err != nil {
			log.Printf("[message handler] response could not be turned in to json %s", req.Method)
			return
		}

		msg.Respond(bs)
	}
}

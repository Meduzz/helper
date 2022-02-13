package starters

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	"github.com/nats-io/nats.go"
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

		if *queue != "" {
			conn.QueueSubscribe(fmt.Sprintf("%s.%s", module, method), *queue, wrapHandler(handler))
		} else {
			conn.Subscribe(fmt.Sprintf("%s.%s", module, method), wrapHandler(handler))
		}

		return nil
	}

	return cmd
}

func wrapHandler(handler wendy.Handler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		req := &wendy.Request{}
		err := json.Unmarshal(msg.Data, req)

		if err != nil {
			log.Printf("[message handler] data could not be parsed to wendy request: %s", msg.Subject)
			return
		}

		res := handler(req)

		bs, err := json.Marshal(res)

		if err != nil {
			log.Printf("[message handler] response could not be turned in to json %s", req.Method)
			return
		}

		msg.Respond(bs)
	}
}

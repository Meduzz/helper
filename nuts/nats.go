package nuts

import (
	"log"
	"os"

	"github.com/nats-io/go-nats"
)

func Connect() (*nats.Conn, error) {
	natsUrl := os.Getenv("NATS_URL")
	natsUser := os.Getenv("NATS_USER")
	natsToken := os.Getenv("NATS_TOKEN")

	if natsUrl == "" {
		log.Println("No NATS_URL, connecting to localhost.")
		natsUrl = "nats://localhost:4222"
	}
	opts := make([]nats.Option, 0)
	if natsUser != "" {
		opts = append(opts, nats.UserInfo(natsUser, os.Getenv("NATS_PASS")))
	} else if natsToken != "" {
		opts = append(opts, nats.Token(natsToken))
	} else {
		log.Println("Neither NATS_USER & NATS_PASS or NATS_TOKEN was provided.")
	}

	return nats.Connect(natsUrl, opts...)
}
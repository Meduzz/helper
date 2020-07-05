package nuts

import (
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
)

func Connect() (*nats.Conn, error) {
	natsUrl := os.Getenv("NATS_URL")
	natsUser := os.Getenv("NATS_USER")
	natsToken := os.Getenv("NATS_TOKEN")

	return Connnect(natsUrl, natsUser, natsToken)
}

func Connnect(natsUrl, natsUser, natsToken string) (*nats.Conn, error) {
	if natsUrl == "" {
		log.Println("No NATS_URL, connecting to localhost.")
		natsUrl = nats.DefaultURL
	}
	opts := make([]nats.Option, 0)
	if natsUser != "" {
		log.Println("NATS_USER set, trying NATS_PASS.")
		opts = append(opts, nats.UserInfo(natsUser, os.Getenv("NATS_PASS")))
	} else if natsToken != "" {
		opts = append(opts, nats.Token(natsToken))
	} else {
		log.Println("Neither NATS_USER & NATS_PASS or NATS_TOKEN was provided.")
	}

	return nats.Connect(natsUrl, opts...)
}

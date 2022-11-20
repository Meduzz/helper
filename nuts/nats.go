package nuts

import (
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
)

// Attempt to connect to nats via data in ENV
func Connect() (*nats.Conn, error) {
	natsUrl := os.Getenv("NATS_URL")
	natsUser := os.Getenv("NATS_USER")
	natsPass := os.Getenv("NATS_PASS")
	natsToken := os.Getenv("NATS_TOKEN")

	opts := make([]nats.Option, 0)

	if natsUrl == "" {
		log.Println("No NATS_URL, connecting to localhost.")
		natsUrl = nats.DefaultURL
	}

	if natsToken != "" {
		opts = append(opts, nats.Token(natsToken))
	} else if natsUser != "" && natsPass != "" {
		opts = append(opts, nats.UserInfo(natsUser, natsPass))
	} else {
		log.Println("Neither NATS_USER & NATS_PASS or NATS_TOKEN was provided.")
	}

	return nats.Connect(natsUrl, opts...)
}

func JsonConnect() (*nats.EncodedConn, error) {
	conn, err := Connect()

	if err != nil {
		return nil, err
	}

	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}

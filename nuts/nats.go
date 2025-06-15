package nuts

import (
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
)

type (
	Config struct {
		Host     string `json:"host"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
)

// Attempt to connect to nats via data in ENV
func Connect() (*nats.Conn, error) {
	natsUrl := os.Getenv("NATS_URL")
	natsUser := os.Getenv("NATS_USER")
	natsPass := os.Getenv("NATS_PASS")
	natsToken := os.Getenv("NATS_TOKEN")

	return Connnect(natsUrl, natsUser, natsPass, natsToken)
}

func Connnect(url, user, pass, token string) (*nats.Conn, error) {
	opts := make([]nats.Option, 0)

	if url == "" {
		log.Println("No NATS_URL, connecting to localhost.")
		url = nats.DefaultURL
	}

	if token != "" {
		opts = append(opts, nats.Token(token))
	} else if user != "" && pass != "" {
		opts = append(opts, nats.UserInfo(user, pass))
	} else {
		log.Println("Neither NATS_USER & NATS_PASS or NATS_TOKEN was provided.")
	}

	return nats.Connect(url, opts...)
}

func ConnectFromConfig(cfg *Config) (*nats.Conn, error) {
	return Connnect(cfg.Host, cfg.Username, cfg.Password, cfg.Token)
}

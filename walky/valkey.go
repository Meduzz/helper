package walky

import (
	"os"
	"strings"

	"github.com/valkey-io/valkey-go"
)

type (
	Config struct {
		Hosts    []string `json:"hosts"`
		Username string   `json:"username,omitempty"`
		Password string   `json:"password,omitempty"`
	}
)

func Connect() (valkey.Client, error) {
	hosts := os.Getenv("CACHE_URL")
	username := os.Getenv("CACHE_USERNAME")
	password := os.Getenv("CACHE_PASSWORD")

	if strings.Contains(hosts, ",") {
		return Connnect(strings.Split(hosts, ","), username, password)
	} else {
		if hosts != "" {
			return Connnect([]string{hosts}, username, password)
		} else {
			return Connnect([]string{"localhost:6379"}, username, password)
		}
	}
}

func Connnect(hosts []string, username, password string) (valkey.Client, error) {
	opts := valkey.ClientOption{}

	opts.InitAddress = hosts

	if username != "" {
		opts.Username = username

		if password != "" {
			opts.Password = password
		}
	}

	return valkey.NewClient(opts)
}

func ConnectFromConfig(cfg *Config) (valkey.Client, error) {
	return Connnect(cfg.Hosts, cfg.Username, cfg.Password)
}

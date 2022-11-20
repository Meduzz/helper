package rudis

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func Connect() *redis.Client {
	address := os.Getenv("REDIS_ADDRESS")
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")

	return Connnect(address, username, password)
}

func Connnect(address, username, password string) *redis.Client {
	if address == "" {
		log.Println("No REDIS_ADDRESS, using localhost:6379")
	}

	if password == "" {
		log.Println("No REDIS_PASSWORD provided")
	}

	if username == "" {
		log.Println("No REDIS_USERNAME provided")
	}

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
	})
}

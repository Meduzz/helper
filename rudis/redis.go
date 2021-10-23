package rudis

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func Connect() *redis.Client {
	address := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")

	return Connnect(address, password)
}

func Connnect(address, password string) *redis.Client {
	if address == "" {
		log.Println("No REDIS_ADDRESS, using localhost:6379")
	}

	if password == "" {
		log.Println("No REDIS_PASSWORD provided")
	}

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
}

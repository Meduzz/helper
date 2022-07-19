package block

import (
	"os"
	"os/signal"
)

//Block will wait for a SIGINT and then execute the provided cleanup function and return any errors it causes.
func Block(cleanup func() error) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return cleanup()
}

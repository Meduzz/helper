package starters

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// Gin - starts a gin webserver based on what's setup in the setup lambda.
func Gin(setup func(*gin.Engine)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a gin server",
	}

	port := cmd.Flags().Int("port", 8080, "port to bind to, defaults to 8080")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		srv := gin.Default()
		setup(srv)

		return srv.Run(fmt.Sprintf(":%d", *port))
	}

	return cmd
}

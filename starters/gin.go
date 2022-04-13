package starters

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// Gin - starts a gin webserver based on what's setup in the setup lambda.
func Gin(port int, setup func(*gin.Engine)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a gin server",
	}

	actualPort := cmd.Flags().Int("port", port, fmt.Sprintf("port to bind to, defaults to %d", port))

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		srv := gin.Default()
		setup(srv)

		return srv.Run(fmt.Sprintf(":%d", *actualPort))
	}

	return cmd
}

package starters

import "github.com/spf13/cobra"

func Generic(handler func(cmd *cobra.Command, args []string) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start the generic starter",
	}

	cmd.RunE = handler

	return cmd
}

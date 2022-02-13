package starters

import "github.com/spf13/cobra"

func Root(version string) *cobra.Command {
	return &cobra.Command{
		Use:     "server",
		Version: version,
	}
}

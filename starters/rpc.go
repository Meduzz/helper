package starters

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/spf13/cobra"
)

func Rpc(setup func(*rpc.RPC)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a rpc module",
	}

	cmd.RunE = func(cdm *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		srv := rpc.NewRpc(conn)

		setup(srv)

		srv.Run()

		return nil
	}

	return cmd
}

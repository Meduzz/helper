package starters

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/encoding"
	"github.com/spf13/cobra"
)

func Rpc(setup func(*cobra.Command, *rpc.RPC)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a rpc module",
	}

	cmd.RunE = func(cdm *cobra.Command, args []string) error {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		srv := rpc.NewRpc(conn, encoding.Json())

		setup(cmd, srv)

		srv.Run()

		return nil
	}

	return cmd
}

package rpc

import (
	"fmt"

	"github.com/Meduzz/helper/service"
	"github.com/Meduzz/rpc"
)

type (
	RpcApi interface {
		Setup(router *rpc.RPC)
	}

	rpcDelegate struct{}
)

var _ service.Delegate = rpcDelegate{}

func (rpcDelegate) Visit(svc service.Service) error {
	api, ok := svc.(RpcApi)

	if ok {
		if server != nil {
			api.Setup(server)
		} else {
			return fmt.Errorf("server not setup")
		}
	}

	return nil
}

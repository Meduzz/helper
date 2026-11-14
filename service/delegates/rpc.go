package delegates

import (
	"github.com/Meduzz/helper/service"
	"github.com/Meduzz/rpc"
)

type (
	WithRPC interface {
		WithRPC(*rpc.RPC)
	}

	rpcDelegate struct {
		router *rpc.RPC
	}
)

var _ service.Delegate = (*rpcDelegate)(nil)

func NewRpcDelegate(router *rpc.RPC) service.Delegate {
	return &rpcDelegate{
		router: router,
	}
}

func (d *rpcDelegate) Visit(svc service.Service) error {
	if api, ok := svc.(WithRPC); ok {
		api.WithRPC(d.router)
	}
	return nil
}

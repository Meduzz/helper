package rpc

import "github.com/Meduzz/rpc"

var (
	server *rpc.RPC
)

func SetRpc(srv *rpc.RPC) {
	server = srv
}

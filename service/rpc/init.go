package rpc

import "github.com/Meduzz/helper/service"

func init() {
	service.AddDelegate(rpcDelegate{})
}

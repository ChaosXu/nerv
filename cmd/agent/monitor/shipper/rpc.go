package shipper

import (
	"github.com/ChaosXu/nerv/lib/env"
	"net/rpc/jsonrpc"
	"log"
)

type RpcShipper struct {
	server string
	cfg    *env.Properties
}

func NewRpcShipper(cfg *env.Properties) Shipper {
	address := cfg.GetMapString("shipper", "server", "3334")
	return &RpcShipper{server:address, cfg:cfg}
}

func (p *RpcShipper) Send(v interface{}) {
	//TBD: client pool
	client, err := jsonrpc.Dial("tcp", p.server)
	if err != nil {
		log.Printf("rpc client dial error. %s", err.Error())
		return
	}
	defer client.Close()

	var out string
	if err := client.Call("MonitorPublisher.Publish", v, &out); err != nil {
		log.Printf("publish error. %s", err.Error())
	}
}


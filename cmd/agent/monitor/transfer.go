package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"net/rpc/jsonrpc"
	"log"
)

//Transfer upload data to the server
type Transfer interface {
	Send(v interface{})
}

type RpcTransfer struct {
	server string
	cfg    *env.Properties
}

func NewRpcTransfer(cfg *env.Properties) Transfer {
	address := cfg.GetMapString("metrics", "server", "3334")
	return &RpcTransfer{server:address, cfg:cfg}
}

func (p *RpcTransfer) Send(v interface{}) {
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


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

type DefaultTransfer struct {
	server string
}

func NewTransfer() Transfer {
	address := env.Config().GetMapString("metrics", "server", "3334")
	return &DefaultTransfer{server:address}
}

func (p *DefaultTransfer) Send(v interface{}) {
	//TBD: client pool
	client, err := jsonrpc.Dial("tcp", p.server)
	if err != nil {
		log.Printf("Push sample error. %s", err.Error())
		return
	}
	defer client.Close()

	var out string
	if err := client.Call("Metrics.Push", v, &out); err != nil {
		log.Printf("Push sample error. %s", err.Error())
	}
}

type LogTransfer struct {

}

func NewLogTransfer() Transfer {
	return &LogTransfer{}
}

func (p *LogTransfer) Send(v interface{}) {
	log.Printf("Send: %+v\n", v)
}

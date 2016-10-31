package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"net/rpc/jsonrpc"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"log"
)

//Transfer upload resource to server
type Transfer interface {
	Send(sample *probe.Sample)
}

type DefaultTransfer struct {
	server string
}

func NewTransfer() Transfer {
	address := env.Config().GetMapString("metrics", "server", "3334")
	return &DefaultTransfer{server:address}
}

func (p *DefaultTransfer) Send(sample *probe.Sample) {
	//TBD: client pool
	client, err := jsonrpc.Dial("tcp", p.server)
	if err != nil {
		log.Printf("Push sample error. %s", err.Error())
		return
	}
	defer client.Close()

	var out string
	if err := client.Call("Metrics.Push", sample, &out); err != nil {
		log.Printf("Push sample error. %s", err.Error())
	} else {
		log.Printf("Push sampleL %s %s %s", sample.Tags["resourceType"], sample.Tags["ip"], sample.Metric, out)
	}
}

type LogTransfer struct {

}

func NewLogTransfer() Transfer {
	return &LogTransfer{}
}

func (p *LogTransfer) Send(sample *probe.Sample) {
	log.Printf("Send sample: %+v\n", sample)
}

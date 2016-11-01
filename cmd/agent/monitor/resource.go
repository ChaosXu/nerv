package monitor

import (
	"fmt"
	"github.com/toolkits/net"
	"github.com/ChaosXu/nerv/lib/debug"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"log"
)

//Resource is the object discovered
type Resource struct {
	Type       string          //resource type
	Address    string          //resource ip
	components []*probe.Sample //resource components
	chSamples  chan *probe.Sample
}

func NewResourceFromSample(sample *probe.Sample) *Resource {
	r := &Resource{
		Type:sample.Tags["resourceType"],
		Address: getLocalAddress(),
		components:[]*probe.Sample{},
		chSamples:make(chan *probe.Sample, 10),
	}
	go r.watchChSamples()
	return r
}
func NewResource(resType string) *Resource {
	//NOTE: Now all resource is discovered in local host,so address is local
	r := &Resource{
		Type:resType,
		Address: getLocalAddress(),
		components:[]*probe.Sample{},
		chSamples:make(chan *probe.Sample, 10),

	}
	go r.watchChSamples()
	return r
}

func getLocalAddress() string {
	if ips, err := net.IntranetIP(); err != nil {
		panic(fmt.Errorf("%s %s", err.Error(), debug.CodeLine()))
	} else {
		return ips[0]
	}
}

func (p *Resource) AddComponent(c *probe.Sample) {
	p.chSamples <- c
}

func (p *Resource) watchChSamples() {
	for {
		c := <-p.chSamples
		p.components = append(p.components, c)
	}
}

func (p *Resource) Key() string {
	return p.Type + "/" + p.Address
}



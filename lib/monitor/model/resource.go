package model

import (
	"fmt"
	"github.com/toolkits/net"
	"github.com/ChaosXu/nerv/lib/debug"
)

//Resource is the object discovered
type Resource struct {
	Type       string          //resource type
	Address    string          //resource ip
	components []*Sample //resource components
	chSamples  chan *Sample
}

func NewResourceFromSample(sample *Sample) *Resource {
	r := &Resource{
		Type:sample.Tags["resourceType"],
		Address: getLocalAddress(),
		components:[]*Sample{},
		chSamples:make(chan *Sample, 10),
	}
	go r.watchChSamples()
	return r
}
func NewResource(resType string) *Resource {
	//NOTE: Now all resource is discovered in local host,so address is local
	r := &Resource{
		Type:resType,
		Address: getLocalAddress(),
		components:[]*Sample{},
		chSamples:make(chan *Sample, 10),

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

func (p *Resource) AddComponent(c *Sample) {
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



package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
)

//Monitor
type Monitor struct {
	discovery *Discovery
	//collector *Collector
	transfer  Transfer
}

func NewMonitor() *Monitor {
	probe := probe.NewProbe()
	transfer := NewLogTransfer()
	discovery := NewDiscovery(probe)
	//collector := NewCollector(nil, probe, transfer)
	return &Monitor{discovery:discovery, transfer:transfer}
}

//Start monitor
func (p *Monitor) Start() error {
	p.startDiscovery()
	//p.collector.Start()
	return nil
}

func (p *Monitor) startDiscovery() {
	path := env.Config().GetMapString("discovery", "path", "../config/discovery")
	templates, err := LoadDiscoveryTemplates(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, template := range templates {
		p.discovery.Add(template)
	}

	go p.discovery.Discover()

	go func() {
		for res := range p.discovery.C {
			p.transfer.Send(res)
		}
	}()
}

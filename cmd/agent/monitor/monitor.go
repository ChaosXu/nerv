package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
)

//Monitor
type Monitor struct {
	discovery *Discovery
	collector *Collector
	transfer  Transfer
}

func NewMonitor() *Monitor {
	probe := probe.NewProbe()
	transfer := NewLogTransfer()
	discovery := NewDiscovery(probe)
	collector := NewCollector(probe, transfer)
	return &Monitor{discovery:discovery, collector:collector, transfer:transfer}
}

//Start monitor
func (p *Monitor) Start() {
	p.startDiscovery()
	p.startCollector()
	go func() {
		for {
			res := <-p.discovery.C
			log.Printf("Push resource: %+v\n", res)
			p.collector.Collect(res)
		}
	}()
}

func (p *Monitor) startDiscovery() {
	path := env.Config().GetMapString("discovery", "path", "../resources/discovery")

	templates, err := LoadDiscoveryTemplates(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, template := range templates {
		p.discovery.Add(template)
	}
	p.discovery.Start()
}

func (p *Monitor) startCollector() {
	path := env.Config().GetMapString("monitor", "path", "../resources/monitor")
	templates, err := LoadMonitorTemplates(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, template := range templates {
		p.collector.Add(template)
	}
	p.collector.Start()
}

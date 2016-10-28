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
}

func NewMonitor() *Monitor {
	probe := probe.NewProbe()
	transfer := NewLogTransfer()
	discovery := NewDiscovery(probe, transfer)
	collector := NewCollector(probe, transfer)
	return &Monitor{discovery:discovery, collector:collector}
}

//Start monitor
func (p *Monitor) Start() error {
	p.startDiscovery()
	p.startCollector()
	return nil
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

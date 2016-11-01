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
	cfg       *env.Properties
}

func NewMonitor(cfg *env.Properties) *Monitor {
	probe := probe.NewProbe(cfg)
	transfer := NewLogTransfer()
	discovery := NewDiscovery(cfg, probe)
	collector := NewCollector(cfg,probe)
	return &Monitor{
		discovery:discovery,
		collector:collector,
		transfer:transfer,
		cfg:cfg,
	}
}

//Start monitor
func (p *Monitor) Start() {
	p.startDiscovery()
	p.startCollector()
	go func() {
		for {
			res := <-p.discovery.C
			p.transfer.Send(res)
			p.collector.Collect(res)
		}
	}()

	go func() {
		for {
			p.transfer.Send(<-p.collector.C)
		}
	}()
}

func (p *Monitor) startDiscovery() {
	path := p.cfg.GetMapString("discovery", "path", "../resources/discovery")

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
	path := p.cfg.GetMapString("monitor", "path", "../resources/monitor")
	templates, err := LoadMonitorTemplates(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, template := range templates {
		p.collector.Add(template)
	}
	p.collector.Start()
}

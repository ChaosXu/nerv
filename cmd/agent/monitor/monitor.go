package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/shipper"
)

//Monitor
type Monitor struct {
	discovery *Discovery
	collector *Collector
	shipper   shipper.Shipper
	cfg       *env.Properties
}

func NewMonitor(cfg *env.Properties) *Monitor {
	probe := probe.NewProbe(cfg)
	discovery := NewDiscovery(cfg, probe)
	collector := NewCollector(cfg, probe)
	return &Monitor{
		discovery:discovery,
		collector:collector,
		shipper:newShipper(cfg),
		cfg:cfg,
	}
}
func newShipper(cfg *env.Properties) shipper.Shipper {
	t:=cfg.GetMapString("shipper","type","rpc")
	switch t {
	case "elasticsearch":
		return shipper.NewElasticsearchShipper(cfg)
	default:
		return shipper.NewRpcShipper(cfg)
	}


}

//Start monitor
func (p *Monitor) Start() {
	p.startDiscovery()
	p.startCollector()
	go func() {
		for {
			res := <-p.discovery.C
			p.shipper.Send(res)
			p.collector.Collect(res)
		}
	}()

	go func() {
		for {
			p.shipper.Send(<-p.collector.C)
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

package monitor

import (
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"github.com/ChaosXu/nerv/lib/monitor/shipper"
	"github.com/ChaosXu/nerv/lib/monitor/shipper/elasticsearch"
	"github.com/ChaosXu/nerv/lib/monitor/shipper/rpc"
	"github.com/ChaosXu/nerv/lib/monitor/model"
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
	shipper := newShipper(cfg)
	return &Monitor{
		discovery:discovery,
		collector:collector,
		shipper:shipper,
		cfg:cfg,
	}
}
func newShipper(cfg *env.Properties) shipper.Shipper {
	t := cfg.GetMapString("shipper", "type", "rpc")
	switch t {
	case "elasticsearch":
		return elasticsearch.NewShipper(cfg)
	default:
		return rpc.NewShipper(cfg)
	}

}

//Start monitor
func (p *Monitor) Start() {
	p.shipper.Init()
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

	templates, err := model.LoadDiscoveryTemplates(path)
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
	templates, err := model.LoadMonitorTemplates(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, template := range templates {
		p.collector.Add(template)
	}
	p.collector.Start()
}

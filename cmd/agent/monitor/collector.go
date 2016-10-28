package monitor

import (
	"log"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
)

//Collector collect all resources's metrics on localhost
type Collector struct {
	templates []*model.MonitorTemplate
	probe     probe.Probe
	transfer  Transfer
}

func NewCollector(probe probe.Probe, transfer Transfer) *Collector {
	return &Collector{probe:probe, transfer:transfer}
}

func (p *Collector) Add(template *model.MonitorTemplate) {
	p.templates = append(p.templates, template)
}

func (p *Collector) Start() error {
	log.Printf("Collector start\n")
	//p.metrics = loadMetrics(env.Config().GetMapString("collector", "path", "../config/metrics"))
	//p.templates = LoadMonitorTemplates(env.Config().GetMapString("collector", "path", "../config/metrics"))
	//
	//period, err := strconv.ParseInt(env.Config().GetMapString("collector", "period", "30"), 10, 0)
	//if err != nil {
	//	return err
	//}
	//p.ticker = time.NewTicker(time.Duration(period) * time.Second)
	//go p.do()

	return nil
}

func (p *Collector) do() {
	//for range p.ticker.C {
	//	go func() {
	//		p.mutex.Lock()
	//		defer p.mutex.Unlock()
	//		if p.doing {
	//			return
	//		}
	//
	//		p.doing = true
	//		for _, metric := range p.metrics {
	//			samples := p.probe.Table(p.ep, metric)
	//			for _, sample := range samples {
	//				p.transfer.Send(sample)
	//			}
	//		}
	//	}()
	//}
}


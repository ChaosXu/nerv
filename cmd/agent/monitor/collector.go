package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/env"
	"strconv"
	"time"
	"sync"
)

//Collector collect all resources's metrics on localhost
type Collector struct {
	templates []*model.MonitorTemplate
	metrics   []*model.Metric
	doing     bool
	mutex     sync.Mutex
	ticker    *time.Ticker
	probe     Probe
	transfer  Transfer
}

func NewCollector(probe Probe, transfer Transfer) *Collector {
	return &Collector{metrics:[]*model.Metric{}, probe:probe, transfer:transfer}
}

func (p *Collector) Start() error {
	p.metrics = loadMetrics(env.Config().GetMapString("collector", "path", "../config/metrics"))
	p.templates = LoadTemplates(env.Config().GetMapString("collector", "path", "../config/metrics"))

	period, err := strconv.ParseInt(env.Config().GetMapString("collector", "period", "30"), 10, 0)
	if err != nil {
		return err
	}
	p.ticker = time.NewTicker(time.Duration(period) * time.Second)
	go p.do()

	return nil
}

func (p *Collector) do() {
	for range p.ticker.C {
		go func() {
			p.mutex.Lock()
			defer p.mutex.Unlock()
			if p.doing {
				return
			}

			p.doing = true
			for _, metric := range p.metrics {
				samples := p.probe.Table(metric)
				for _, sample := range samples {
					p.transfer.Send(sample)
				}
			}
		}()
	}
}


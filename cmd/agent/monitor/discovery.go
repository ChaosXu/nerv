package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/env"
	"strconv"
	"time"
	"sync"
)

//Discovery search localhost to find resource
type Discovery struct {
	metrics   []*model.Metric
	doing     bool
	mutex     sync.Mutex
	ticker    *time.Ticker
	collector Collector
	transfer  Transfer
}

func NewDiscovery(collector Collector, transfer Transfer) *Discovery {
	return &Discovery{metrics:[]*model.Metric{}, collector:collector, transfer:transfer}
}

func (p *Discovery) Start() error {
	period, err := strconv.ParseInt(env.Config().GetMapString("discovery", "period", "30"), 10, 0)
	if err != nil {
		return err
	}

	p.ticker = time.NewTicker(time.Duration(period) * time.Second)

	go p.do()
	return nil
}

func (p *Discovery) do() {
	for range p.ticker.C {
		go func() {
			p.mutex.Lock()
			defer p.mutex.Unlock()
			if p.doing {
				return
			}

			p.doing = true
			for _, metric := range p.metrics {
				samples := p.collector.Table(metric)
				for _, sample := range samples {
					p.transfer.Send(sample)
				}
			}
		}()
	}
}

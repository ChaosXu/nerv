package monitor

import (
	"log"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
)

//Collector collect all resources's metrics on localhost
type Collector struct {
	watchers map[string]map[int64]*Watcher
	probe    probe.Probe
	C        chan *model.Sample
}

func NewCollector(probe probe.Probe) *Collector {
	return &Collector{
		watchers:map[string]map[int64]*Watcher{},
		probe:probe,
	}
}

func (p *Collector) Add(template *model.MonitorTemplate) {
	periods := p.watchers[template.ResourceType]
	if periods == nil {
		periods = map[int64]*Watcher{}
		p.watchers[template.ResourceType] = periods
	}

	for _, v := range template.Items {
		watcher := periods[v.Period]
		if watcher == nil {
			watcher = NewWatcher(template.ResourceType, v.Period, p.probe)
			periods[v.Period] = watcher
		}
		watcher.AddItem(v)
	}
}

func (p *Collector) Start() {
	log.Printf("Collector start\n")
	p.C = make(chan *model.Sample, 1000)
	for _, periods := range p.watchers {
		for _, w := range periods {
			w.Start(p.C)
		}
	}
}

func (p *Collector) Stop() {
	if p.C != nil {
		for _, periods := range p.watchers {
			for _, w := range periods {
				w.Stop()
			}
		}
		close(p.C)
		p.C = nil
	}

}

//Collect resource's metrics if it match a monitor template
func (p *Collector) Collect(resource *model.Resource) {
	periods := p.watchers[resource.Type]
	if periods == nil {
		return
	}

	for _, watcher := range periods {
		watcher.AddResource(resource)
	}
}


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
	transfer Transfer
}

func NewCollector(probe probe.Probe, transfer Transfer) *Collector {
	return &Collector{
		watchers:map[string]map[int64]*Watcher{},
		probe:probe,
		transfer:transfer,
	}
}

func (p *Collector) Add(template *model.MonitorTemplate) {
	periods := map[int64]*Watcher{}
	p.watchers[template.Name] = periods
	for _, v := range template.Items {
		watcher := periods[v.Period]
		if watcher == nil {
			watcher = NewWatcher(template.ResourceType)
			periods[v.Period] = watcher
		}
		watcher.AddItem(v)
		log.Printf("watcher add item: %s %s %ds\n", template.ResourceType, v.Metric, v.Period)
	}
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

//Collect resource's metrics if it match a monitor template
func (p *Collector) Collect(resource *Resource) {
	//TBD: watch once if there is more than one period for one metric
	periods := p.watchers[resource.Type]
	if periods == nil {
		return
	}

	for _, watcher := range periods {
		watcher.Watch(resource)
	}
}


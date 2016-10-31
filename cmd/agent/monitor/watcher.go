package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"time"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"log"
	"github.com/ChaosXu/nerv/lib/debug"
)

//Watcher collect sample of the metric
type Watcher struct {
	resourceType string
	period       int64
	probe        probe.Probe
	cancel       chan struct{}
	items        []model.MonitorItem
	resources    []*Resource
}

//NewWatcher create a watcher
func NewWatcher(resType string, period int64, probe probe.Probe) *Watcher {
	return &Watcher{
		resourceType:resType,
		period:period,
		probe:probe,
	}
}

//AddItem add a template for metric
func (p *Watcher) AddItem(item model.MonitorItem) {
	log.Printf("Watcher.AddItem %s %s %ds\n", p.resourceType, item.Metric, item.Period)
	p.items = append(p.items, item)
}

//Watch the resource through monitor items that has been added by AddItem
func (p *Watcher) AddResource(res *Resource) {
	log.Printf("Watcher.AddResoruce %s %d %s\n", res.Type, p.period, debug.CodeLine())
	p.resources = append(p.resources, res)
}

//Start watcher to collect items periodically
func (p *Watcher) Start(out chan <- *probe.Sample) {
	log.Printf("Watcher.Start %s %d\n", p.resourceType, p.period)
	p.cancel = make(chan struct{})

	go func() {
		timer := time.NewTicker(time.Duration(p.period) * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				p.read(out)
			case <-p.cancel:
				return
			}
		}
	}()
}

func (p *Watcher) Stop() {
	log.Printf("Watcher.Stop %s %d\n", p.resourceType, p.period)
	if p.cancel != nil {
		close(p.cancel)
		p.cancel = nil
	}
}

func (p *Watcher) read(out chan <-*probe.Sample) {
	log.Printf("Watcher.read %s %d %d %d\n", p.resourceType, p.period, len(p.resources), len(p.items))
	for _, res := range p.resources {
		p.readItem(res, out)
	}
}

func (p *Watcher) readItem(res *Resource, out chan <-*probe.Sample) {
	for _, item := range p.items {
		log.Printf("Watcher.readItem. %s %s %s \n", p.resourceType, item.Metric, debug.CodeLine())
		if metric, err := loadMetric(p.resourceType, item.Metric); err != nil {
			log.Printf("Watcher.readerItem error. %s %s %s %s\n", p.resourceType, item.Metric, err.Error(), debug.CodeLine())
		} else {
			switch metric.Type {
			case model.MetricTypeStruct:
				if sample, err := p.probe.Row(metric, p.metricArgs(item)); err != nil {
					log.Printf("Watcher.readerItem error. %s %s %s %s\n", p.resourceType, item.Metric, err.Error(), debug.CodeLine())
				} else {
					out <- sample
					log.Printf("Watcher.readerItem complete. %s %s %s\n", p.resourceType, item.Metric, debug.CodeLine())
				}
			}
		}
	}
}

func (p *Watcher) metricArgs(item model.MonitorItem) map[string]string {
	return map[string]string{}
}

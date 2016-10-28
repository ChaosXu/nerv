package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"github.com/ChaosXu/nerv/lib/debug"
)

//Resource
type Resource struct {

}

//Discovery search localhost to find resource
type Discovery struct {
	host     *model.DiscoveryTemplate
	services map[string]*model.DiscoveryTemplate
	probe    probe.Probe
	C        chan *probe.Sample
}

func NewDiscovery(p probe.Probe) *Discovery {
	return &Discovery{services:map[string]*model.DiscoveryTemplate{}, probe:p, C:make(chan *probe.Sample, 50)}
}

func (p *Discovery) Add(template *model.DiscoveryTemplate) {
	if template.Host {
		p.host = template
	} else {
		p.services[template.ResourceType] = template
	}
}

func (p *Discovery) Discover() {
	log.Printf("Discover: %s %s\n", p.host.ResourceType, p.host.Name)

	template := p.host
	for _, item := range template.Items {
		metric, err := readMetric(template.ResourceType, item.Metric)
		if err != nil {
			log.Printf("Discover error. %s %s %s %s\n", template.ResourceType, item.Metric, err.Error(), debug.CodeLine())
			continue
		}

		go func(template *model.DiscoveryTemplate, item model.DiscoveryItem, metric *model.Metric) {
			samples, err := p.probe.Table(metric, nil)
			if err != nil {
				log.Printf("Discover error. %s %s %s\n", template.ResourceType, metric.Name, err.Error(), debug.CodeLine())
				return
			}

			if item.Service != "" {
				for _, sample := range samples {
					p.discoverService(item.Service, sample)
				}
			} else {
				for _, sample := range samples {
					p.C <- sample
				}
			}
		}(template, item, metric)

	}
}

func (p *Discovery) discoverService(resourceType string, v *probe.Sample) {
	template := p.services[resourceType]
	if template == nil {
		log.Printf("discoverService: no discovery template for %s %s\n", resourceType, debug.CodeLine())
		return
	}
	log.Printf("discoverService: %s %s\n", resourceType, template.Name)

	for _, item := range template.Items {
		metric, err := readMetric(template.ResourceType, item.Metric)
		if err != nil {
			log.Printf("discoverService error. %s %s %s\n", template.ResourceType, item.Metric, err.Error(), debug.CodeLine())
			continue
		}

		go func(template *model.DiscoveryTemplate, metric *model.Metric) {
			samples, err := p.probe.Table(metric, nil)
			if err != nil {
				log.Printf("discoverService error. %s %s %s\n", template.ResourceType, metric.Name, err.Error(), debug.CodeLine())
				return
			}

			for _, sample := range samples {
				p.C <- sample
			}
		}(template, metric)
	}
}




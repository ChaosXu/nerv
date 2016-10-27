package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
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
	if template.IsService {
		p.services[template.ResourceType] = template
	} else {
		p.host = template
	}
}

func (p *Discovery) Discover() {
	log.Printf("Discovery: %s %s\n", p.host.ResourceType, p.host.Name)

	template := p.host
	for _, item := range template.Items {
		metric, err := readMetric(template.ResourceType, item.Metric)
		if err != nil {
			log.Printf("ReadMetric error. %s.%s %s\n", template.ResourceType, item.Metric, err.Error())
			continue
		}

		go func(template *model.DiscoveryTemplate, metric *model.Metric) {
			samples, err := p.probe.Table(metric, nil)
			if err != nil {
				log.Printf("Probe error. %s.%s\n", template.ResourceType, metric.Name, err.Error())
				return
			}

			if metric.IsService {
				for _, sample := range samples {
					p.discoverApplication(sample)
				}
			} else {
				for _, sample := range samples {
					p.C <- sample
				}
			}
		}(template, metric)

	}
}

func (p *Discovery) discoverApplication(v *probe.Sample) {
	resourceType := v.Tags["resourceType"]
	template := p.services[resourceType]
	if template == nil {
		log.Printf("Discovery: Don't find service template. %s\n", resourceType)
		return
	}
	log.Printf("Discovery: %s %s\n", resourceType, template.Name)

	for _, item := range template.Items {
		metric, err := readMetric(template.ResourceType, item.Metric)
		if err != nil {
			log.Printf("ReadMetric error. %s.%s\n", template.ResourceType, item.Metric, err.Error())
			continue
		}

		go func(template *model.DiscoveryTemplate, metric *model.Metric) {
			samples, err := p.probe.Table(metric, nil)
			if err != nil {
				log.Printf("Probe error. %s.%s\n", template.ResourceType, metric.Name, err.Error())
				return
			}

			for _, sample := range samples {
				p.C <- sample
			}
		}(template, metric)
	}
}




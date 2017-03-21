package monitor

import (
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"github.com/ChaosXu/nerv/lib/debug"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"time"
)

//Discovery search localhost to find resource
type Discovery struct {
	cfg          *env.Properties
	host         *model.DiscoveryTemplate
	services     map[string]*model.DiscoveryTemplate
	probe        probe.Probe
	C            chan *model.Resource
	stopDiscover chan bool
}

func NewDiscovery(cfg *env.Properties, p probe.Probe) *Discovery {
	return &Discovery{
		cfg:          cfg,
		services:     map[string]*model.DiscoveryTemplate{},
		probe:        p,
		C:            make(chan *model.Resource, 1000),
		stopDiscover: make(chan bool, 1),
	}
}

func (p *Discovery) Add(template *model.DiscoveryTemplate) {
	if template.Host {
		p.host = template
	} else {
		p.services[template.ResourceType] = template
	}
}

func (p *Discovery) Start() {
	template := p.host
	period := template.Period
	if period == 0 {
		period = 30
	}
	log.Printf("Discovery start %s %d\n", template.ResourceType, period)

	go func() {
		p.discoverHost()

		ticker := time.NewTicker(time.Duration(period) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.discoverHost()
			case <-p.stopDiscover:
				log.Println("Discovery strop discover")
				return
			}
		}
	}()
}

//func (p *Discovery) Stop() {
//	log.Printf("Discovery start %s %d\n", p.host.ResourceType)
//
//	close(p.stopDiscover)
//	close(p.C)
//}

func (p *Discovery) discoverHost() {
	log.Printf("Discover: %s %s\n", p.host.ResourceType, p.host.Name)

	localhost := p.newLocalHost()
	p.C <- localhost
	template := p.host
	for _, item := range template.Items {
		metric, err := model.LoadMetric(p.cfg, template.ResourceType, item.Metric)
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
					sample.Tags["resourceType"] = item.Service
					service := model.NewResourceFromSample(sample)
					p.C <- service
					p.discoverService(service, item.Service, sample)
				}
			} else {
				for _, sample := range samples {
					localhost.AddComponent(sample)
				}
			}
		}(template, item, metric)
	}
}

func (p *Discovery) discoverService(service *model.Resource, resourceType string, v *model.Sample) {
	template := p.services[resourceType]
	if template == nil {
		log.Printf("discoverService: no discovery template for %s %s\n", resourceType, debug.CodeLine())
		return
	}
	log.Printf("discoverService: %s %s\n", resourceType, template.Name)

	for _, item := range template.Items {
		metric, err := model.LoadMetric(p.cfg, template.ResourceType, item.Metric)
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
				service.AddComponent(sample)
			}
		}(template, metric)
	}
}

func (p *Discovery) newLocalHost() *model.Resource {
	return model.NewResource("/host/Linux")
}

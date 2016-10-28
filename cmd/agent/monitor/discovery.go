package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/monitor/probe"
	"github.com/ChaosXu/nerv/lib/debug"
	"time"
)

//Resource
type Resource struct {

}

//Discovery search localhost to find resource
type Discovery struct {
	host         *model.DiscoveryTemplate
	services     map[string]*model.DiscoveryTemplate
	probe        probe.Probe
	transfer     Transfer
	c            chan *probe.Sample
	stopDiscover chan bool
	stopTransfer chan bool
}

func NewDiscovery(p probe.Probe, transfer Transfer) *Discovery {
	return &Discovery{
		services:map[string]*model.DiscoveryTemplate{},
		probe:p, transfer:transfer,
		c:make(chan *probe.Sample, 50),
		stopDiscover:make(chan bool, 1),
		stopTransfer:make(chan bool, 1),
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
		ticker := time.NewTicker(time.Duration(period) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.discover()
			case <-p.stopDiscover:
				log.Println("Discovery strop discover")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case sample := <-p.c:
				p.transfer.Send(sample)
			case <-p.stopTransfer:
				log.Println("Discovery strop transfer")
				return
			}
		}
	}()
}

func (p *Discovery) Stop() {
	log.Printf("Discovery start %s %d\n", p.host.ResourceType)

	close(p.stopDiscover)
	close(p.stopTransfer)
}

func (p *Discovery) discover() {
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
					p.c <- sample
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
				p.c <- sample
			}
		}(template, metric)
	}
}




package monitor

import "github.com/ChaosXu/nerv/cmd/agent/monitor/probe"

//Resource is the object discovered
type Resource struct {
	Type       string
	components []*probe.Sample
	chSamples  chan *probe.Sample
}

func NewResourceFromSample(sample *probe.Sample) *Resource {
	r := &Resource{Type:sample.Tags["resourceType"], components:[]*probe.Sample{}, chSamples:make(chan *probe.Sample, 10)}
	go r.watchChSamples()
	return r
}
func NewResource(resType string) *Resource {
	r := &Resource{Type:resType, components:[]*probe.Sample{}, chSamples:make(chan *probe.Sample, 10)}
	go r.watchChSamples()
	return r
}

func (p *Resource) AddComponent(c *probe.Sample) {
	p.chSamples <- c
}

func (p *Resource) watchChSamples() {
	for {
		c := <-p.chSamples
		p.components = append(p.components, c)
	}
}



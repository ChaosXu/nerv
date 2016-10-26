package monitor

import "github.com/ChaosXu/nerv/lib/monitor/model"

//Endpoint is a interface to access resource
type Endpoint struct {

}

//Sample is the data collected
type Sample struct {

}

//Probe collects the data of metrics
type Probe interface {
	Table(ep *Endpoint, metric *model.Metric) []*Sample
}

type DefaultProbe struct {

}

func NewProbe() Probe {
	return &DefaultProbe{}
}
//Table return table data
func (p *DefaultProbe) Table(ep *Endpoint, metric *model.Metric) []*Sample {
	return nil
}

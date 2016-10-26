package monitor

import "github.com/ChaosXu/nerv/lib/monitor/model"

//Sample is the data collected
type Sample struct {

}

//Probe collects the data of metrics
type Probe interface {
	Table(metric *model.Metric) []*Sample
}

type DefaultProbe struct{

}

func NewCollector() Probe {
	return &DefaultProbe{}
}
//Table return table data
func (p *DefaultProbe) Table(metric *model.Metric) []*Sample {
	return nil
}

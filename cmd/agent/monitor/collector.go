package monitor

import "github.com/ChaosXu/nerv/lib/monitor/model"

//Sample is the data collected
type Sample struct {

}

//Collector collects the data of metrics
type Collector interface {
	Table(metric *model.Metric) []*Sample
}

type DefaultCollector struct{

}

func NewCollector() Collector {
	return &DefaultCollector{}
}
//Table return table data
func (p *DefaultCollector) Table(metric *model.Metric) []*Sample {
	return nil
}

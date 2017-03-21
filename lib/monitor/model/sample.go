package model

import "time"

//Sample is the data collected
type Sample struct {
	Metric    string                 //metric name
	Values    map[string]interface{} //values of metric fields
	Tags      map[string]string      //every sample has default tags: resourceType,ip:
	Timestamp int64                  //utc time
}

func NewSample(metric string, values map[string]interface{}, resourceType string) *Sample {
	tags := map[string]string{
		"resourceType": resourceType,
	}
	return &Sample{Metric: metric, Values: values, Tags: tags, Timestamp: time.Now().Unix()}
}

func (p *Sample) Merge(other *Sample) {
	for k, v := range other.Values {
		p.Values[k] = v
	}

	for k, v := range other.Tags {
		p.Values[k] = v
	}
}

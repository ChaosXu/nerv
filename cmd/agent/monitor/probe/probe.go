package probe

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"github.com/ChaosXu/nerv/lib/debug"
	"time"
)

//Sample is the data collected
type Sample struct {
	Metric    string                 //metric name
	Values    map[string]interface{} //values of metric fields
	Tags      map[string]string      //every sample has default tags: resourceType,ip:
	Timestamp int64                  //utc time
}

func NewSample(metric string, values map[string]interface{}, resourceType string) *Sample {
	tags := map[string]string{
		"resourceType":resourceType,
	}
	return &Sample{Metric:metric, Values:values, Tags:tags, Timestamp:time.Now().Unix()}
}

func (p *Sample) Merge(other *Sample) {
	for k, v := range other.Values {
		p.Values[k] = v
	}

	for k, v := range other.Tags {
		p.Values[k] = v
	}
}

//Probe collects the data of metrics
type Probe interface {
	Table(metric *model.Metric, args map[string]string) ([]*Sample, error)
}

type DefaultProbe struct {
	probes map[model.ProbeType]Probe
}

func NewProbe() Probe {
	probes := map[model.ProbeType]Probe{
		model.ProbeTypeShell: NewShellProbe(),
	}
	return &DefaultProbe{probes:probes}
}

//Table return table data
func (p *DefaultProbe) Table(metric *model.Metric, args map[string]string) ([]*Sample, error) {
	log.Printf("DefaultProbe.Table %s %s %s", metric.ResourceType, metric.Name, debug.CodeLine())
	chSamples := []chan []*Sample{}
	for _, probe := range p.probes {
		ch := make(chan []*Sample, 1)
		chSamples = append(chSamples, ch)
		go func(probe Probe, ch chan <- []*Sample) {
			samples, err := probe.Table(metric, args)
			if err != nil {
				log.Printf("Probe.Table error. %s", err.Error())
				ch <- []*Sample{}
			} else {
				ch <- samples
			}
		}(probe, ch)
	}

	key := metric.Key()
	sampleMap := map[interface{}]*Sample{}

	for _, ch := range chSamples {
		ss := <-ch
		for _, s := range ss {
			sample := sampleMap[s.Values[key]]
			if sample == nil {
				sample = s
				sampleMap[s.Values[key]] = sample
			} else {
				sample.Merge(s)
			}
		}
	}

	samples := make([]*Sample, len(sampleMap))
	i := 0
	for _, s := range sampleMap {
		samples[i] = s
		i++
	}
	return samples, nil
}

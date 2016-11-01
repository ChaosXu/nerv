package probe

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"log"
	"github.com/ChaosXu/nerv/lib/debug"
	"github.com/ChaosXu/nerv/lib/env"
)

//Probe collects the data of metrics
type Probe interface {
	Table(metric *model.Metric, args map[string]string) ([]*model.Sample, error)
	Row(metric *model.Metric, args map[string]string) (*model.Sample, error)
}

type DefaultProbe struct {
	probes map[model.ProbeType]Probe
}

func NewProbe(cfg *env.Properties) Probe {
	probes := map[model.ProbeType]Probe{
		model.ProbeTypeShell: NewShellProbe(cfg),
	}
	return &DefaultProbe{probes:probes}
}

//Table return table data
func (p *DefaultProbe) Table(metric *model.Metric, args map[string]string) ([]*model.Sample, error) {
	log.Printf("DefaultProbe.Table %s %s %s", metric.ResourceType, metric.Name, debug.CodeLine())
	chSamples := []chan []*model.Sample{}
	for _, probe := range p.probes {
		ch := make(chan []*model.Sample, 1)
		chSamples = append(chSamples, ch)
		go func(probe Probe, ch chan <- []*model.Sample) {
			samples, err := probe.Table(metric, args)
			if err != nil {
				log.Printf("Probe.Table error. %s", err.Error())
				ch <- []*model.Sample{}
			} else {
				ch <- samples
			}
		}(probe, ch)
	}

	key := metric.Key()
	sampleMap := map[interface{}]*model.Sample{}

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

	samples := make([]*model.Sample, len(sampleMap))
	i := 0
	for _, s := range sampleMap {
		samples[i] = s
		i++
	}
	return samples, nil
}

func (p *DefaultProbe) Row(metric *model.Metric, args map[string]string) (*model.Sample, error) {
	log.Printf("DefaultProbe.Row %s %s %s", metric.ResourceType, metric.Name, debug.CodeLine())

	chSamples := []chan *model.Sample{}
	for _, probe := range p.probes {
		ch := make(chan *model.Sample, 1)
		chSamples = append(chSamples, ch)
		go func(probe Probe, ch chan <- *model.Sample) {
			sample, err := probe.Row(metric, args)
			if err != nil {
				log.Printf("Probe.Row error. %s", err.Error())
				ch <- &model.Sample{}
			} else {
				ch <- sample
			}
		}(probe, ch)
	}

	var sample *model.Sample
	for _, ch := range chSamples {
		s := <-ch
		if sample == nil {
			sample = s
		} else {
			sample.Merge(s)
		}
	}

	return sample, nil
}

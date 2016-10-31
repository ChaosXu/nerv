package probe

import (
	"log"
	"github.com/ChaosXu/nerv/lib/debug"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"fmt"
	"os/exec"
	"encoding/json"
	"github.com/ChaosXu/nerv/lib/env"
)

type ShellProbe struct {

}

func NewShellProbe() Probe {
	return &ShellProbe{}
}

func (p *ShellProbe) Table(metric *model.Metric, args map[string]string) ([]*Sample, error) {
	log.Printf("ShellProbe.Table %s %s %s", metric.ResourceType, metric.Name, debug.CodeLine())
	log.Printf("%+v", metric)
	chs := map[string]chan []*Sample{}
	for _, field := range metric.Fields {
		if field.Probe.Type == model.ProbeTypeShell {
			ch := chs[field.Probe.Provider]
			if ch != nil {
				continue
			}

			ch = make(chan []*Sample, 1)
			chs[field.Probe.Provider] = ch
			go func(field model.MetricField, ch chan []*Sample) {
				if res, err := p.exec(field.Probe.Provider, args); err != nil {
					log.Printf("ShellProbe.Table error %s %s %s %s %s", err.Error(), metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())
					ch <- []*Sample{}
				} else {
					log.Printf("ShellProbe.Table %s %s %s %s %s", res, metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())

					switch metric.Type{

					case model.MetricTypeTable:
						v := []map[string]interface{}{}
						if err := json.Unmarshal([]byte(res), &v); err != nil {
							log.Printf("ShellProbe.Table error %s %s %s %s %s", err.Error(), metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())
							ch <- []*Sample{}
						} else {
							samples := []*Sample{}
							for _, item := range v {
								sample := NewSample(metric.Name, item, metric.ResourceType)
								samples = append(samples, sample)
							}
							ch <- samples
						}
					}
				}
			}(field, ch)
		}
	}

	samples := []*Sample{}
	for _, ch := range chs {
		ss := <-ch
		for _, s := range ss {
			samples = append(samples, s)
		}
	}

	return samples, nil
}

func (p *ShellProbe) Row(metric *model.Metric, args map[string]string) (*Sample, error) {
	log.Printf("ShellProbe.Row %s %s %s", metric.ResourceType, metric.Name, debug.CodeLine())
	log.Printf("%+v", metric)
	chs := map[string]chan *Sample{}
	for _, field := range metric.Fields {
		if field.Probe.Type == model.ProbeTypeShell {
			ch := chs[field.Probe.Provider]
			if ch == nil {
				continue
			}
			//read once
			ch = make(chan *Sample, 1)
			chs[field.Probe.Provider] = ch
			go func(field model.MetricField, ch chan *Sample) {
				if res, err := p.exec(field.Probe.Provider, args); err != nil {
					log.Printf("ShellProbe.Row error %s %s %s %s %s", err.Error(), metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())
					ch <- nil
				} else {
					log.Printf("ShellProbe.Row %s %s %s %s %s", res, metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())

					switch metric.Type{
					case model.MetricTypeStruct:
						v := map[string]interface{}{}
						if err := json.Unmarshal([]byte(res), &v); err != nil {
							log.Printf("ShellProbe.Row error %s %s %s %s %s", err.Error(), metric.ResourceType, metric.Name, field.Probe.Provider, debug.CodeLine())
							ch <- nil
						} else {
							ch <- NewSample(metric.Name, v, metric.ResourceType)
						}
					}
				}
			}(field, ch)
		}
	}

	var sample *Sample
	for _, ch := range chs {
		s := <-ch
		if sample == nil {
			sample = s
		} else {
			sample.Merge(s)
		}
	}

	return sample, nil
}

func (p *ShellProbe) exec(file string, args map[string]string) (string, error) {
	log.Printf("ShellProbe.exec %s %s", file, debug.CodeLine())
	root := env.Config().GetMapString("scripts", "path", "../config/scripts")
	export := ""
	for k, v := range args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}

	shell := "export LC_TIME=POSIX " + export + " && " + root + file
	log.Println(shell)

	out, err := exec.Command("/bin/bash", "-c", shell).Output()
	if err != nil {
		return "", err
	}
	s := string(out)
	return s, nil
}

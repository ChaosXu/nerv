package resource

import "github.com/chaosxu/nerv/lib/log"

func init() {
	Models["Agent"] = &Agent{}
}

type Agent struct {

}

func (p *Agent) Create() error {
	log.LogCodeLine()
	return nil
}

func (p *Agent) Delete() error {
	return nil
}

func (p *Agent) Start() error {
	return nil
}

func (p *Agent) Stop() error {
	return nil
}

func (p *Agent) Setup() error {
	return nil
}
package resource

import "github.com/chaosxu/nerv/lib/log"

func init() {
	Models["Host"] = &Host{}
}

type Host struct {

}

func (p *Host) Create() error {
	log.LogCodeLine()
	return nil
}

func (p *Host) Delete() error {
	return nil
}

func (p *Host) Start() error {
	return nil
}

func (p *Host) Stop() error {
	return nil
}

func (p *Host) Setup() error {
	return nil
}
package resource

import "github.com/chaosxu/nerv/lib/log"

func init() {
	Models["Agent"] = &Agent{}
}

//Agent is a proxy for control the process instance on the host
type Agent struct {

}

//Create a process instance on the host
func (p *Agent) Create() error {
	log.LogCodeLine()
	return nil
}

//Delete the process instance on the host
func (p *Agent) Delete() error {
	return nil
}

//Start the process instance on the host
func (p *Agent) Start() error {
	return nil
}

//Stop the process instance on the host
func (p *Agent) Stop() error {
	return nil
}

//Setup the process instance. eg. config
func (p *Agent) Setup() error {
	return nil
}
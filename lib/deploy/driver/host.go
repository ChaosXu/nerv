package driver

import "github.com/ChaosXu/nerv/lib/log"

func init() {
	Models["Host"] = &Host{}
}

//Host is a proxy for control the real host
type Host struct {

}

//Create a proxy in the resource pool
func (p *Host) Create() error {
	log.LogCodeLine()
	return nil
}

//Delete a proxy from the resource pool
func (p *Host) Delete() error {
	return nil
}

//Start the proxy
func (p *Host) Start() error {
	return nil
}

//Stop the proxy
func (p *Host) Stop() error {
	return nil
}

//Setup the proxy
func (p *Host) Setup() error {
	return nil
}
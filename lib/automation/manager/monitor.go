package manager

import "github.com/ChaosXu/nerv/lib/automation/model/topology"

// Monitor manage the task of topology monitoring
type Monitor struct{

}

// Start monitoring for topology
func (p *Monitor) Start(topo *topology.Topology) error {
	return nil
}

// Stop monitoring
func (p *Monitor) Stop(topo *topology.Topology) error {
	return nil
}

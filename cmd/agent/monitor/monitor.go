package monitor

//Monitor
type Monitor struct {
	discovery *Discovery
	collector *Collector
}

func NewMonitor() *Monitor {
	probe := NewProbe()
	transfer := NewTransfer()
	discovery := NewDiscovery(nil,probe, transfer)
	collector := NewCollector(nil,probe, transfer)
	return &Monitor{discovery:discovery, collector:collector}
}

//Start monitor
func (p *Monitor) Start() error {
	p.discovery.Start()
	p.collector.Start()
	return nil
}

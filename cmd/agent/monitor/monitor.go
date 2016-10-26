package monitor

//Monitor
type Monitor struct {
	Discovery *Discovery
}

func NewMonitor() *Monitor {
	collector := NewCollector()
	transfer := NewTransfer()
	discovery := NewDiscovery(collector, transfer)
	return &Monitor{Discovery:discovery}
}

//Start monitor
func (p *Monitor) Start() error {
	p.Discovery.Start()
	return nil
}

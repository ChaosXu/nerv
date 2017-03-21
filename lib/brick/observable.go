package brick

// EventHandler process the event
type EventHandler interface {
	// Handle process the event
	Handle(event string, data interface{})
}

type Observable interface {
	// Register a handler on the event
	On(event string, handler EventHandler)

	// Remove the handler on the event
	Off(event string, handler EventHandler)
}

// Notify events
type Notify interface {
	// Emmit an event
	Emmit(event string, data interface{})
}

type Trigger struct {
	events map[string][]EventHandler
}

func (p *Trigger) On(event string, handler EventHandler) {
	if p.events == nil {
		p.events = map[string][]EventHandler{}
	}
	hs := p.events[event]
	if hs == nil {
		hs = []EventHandler{}
	}
	hs = append(hs, handler)
	p.events[event] = hs
}

func (p *Trigger) Off(event string, handler EventHandler) {
	if p.events == nil {
		return
	}
	hs := p.events[event]
	if hs == nil {
		return
	}
	for i, h := range hs {
		if h != handler {
			hs = append(hs[:i], hs[i+1:]...)
			break
		}
	}
	p.events[event] = hs
}

func (p *Trigger) Emmit(event string, data interface{}) {
	if p.events == nil {
		return
	}
	hs := p.events[event]
	if hs == nil {
		return
	}
	for _, h := range hs {
		h.Handle(event, data)
	}
}

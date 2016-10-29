package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
)

//Watcher collect sample of the metric
type Watcher struct {
	resType string
	items   []model.MonitorItem
}

//NewWatcher create a watcher
func NewWatcher(resType string) *Watcher {
	return &Watcher{resType:resType}
}

//AddItem add a template for metric
func (p *Watcher) AddItem(item model.MonitorItem) {
	p.items = append(p.items, item)
}

//Watch the resource if it match a monitor item.if matched return true.
func (p *Watcher) Watch(res *Resource) bool {
	//for _, item := range p.items {
	//
	//}
	return false
}

package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/log"
)


//Node is element of topology
type Node struct {
	gorm.Model
	Status
	TopologyID int        `gorm:"index"` //Foreign key of the topology
	Name       string                    //node name
	Template   string                    //template name
	Address    string                    //address of node.
	Links      []*Link
}

func init() {
	db.Models["Node"] = nodeDesc()
}

func nodeDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Node{},
		New: func() interface{} {
			return &Node{}
		},
		NewSlice:func() interface{} {
			return &[]Node{}
		},
	}
}

// link the source node to the target node
func (p *Node) Link(depType string, target string) {
	if p.Links == nil {
		p.Links = []*Link{}
	}
	p.Links = append(p.Links, &Link{Type:depType, Source:p.Name, Target:target})
}

// findLinksByType return all links of depType
func (p *Node) FindLinksByType(depType string) []*Link {
	links := []*Link{}
	for _, link := range p.Links {
		if link.Type == depType {
			links = append(links, link)
		}
	}
	return links
}

// Execute operation
func (p *Node) Execute(operation string, nodeTemplate *NodeTemplate) error {
	log.LogCodeLine()

	p.RunStatus = RunStatusGreen

	var err error = nil
	class := Class{}
	if err = db.DB.Where("name=?", nodeTemplate.Type).Preload("Operations").First(&class).Error; err != nil {
		p.RunStatus = RunStatusRed
		p.Error = err.Error()
		db.DB.Save(p)
		return err
	}

	if err = class.Invoke(operation, p, nodeTemplate); err != nil {
		p.RunStatus = RunStatusRed
		p.Error = err.Error()
		db.DB.Save(p)
		return err
	}

	return nil
}

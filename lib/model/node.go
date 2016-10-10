package model

import (
	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/db"
)

type NodeStatus int

const (
	NodeStatusNew NodeStatus = iota //when new
	NodeStatusGreen        //all element ok
	NodeStatusYellow    //some ok,some failed
	NodeStatusRed        //all element failed
)

//Node is element of topology
type Node struct {
	gorm.Model
	TopologyID int        `gorm:"index"` //Foreign key of the topology
	Name       string                    //node name
	Template   string                    //template name
	Links      []*Link
	Status     NodeStatus                //status of node
										 //Error    error      //error during the node process
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

//Link the source node to the target node
func (p *Node) link(depType string, target string) {
	if p.Links == nil {
		p.Links = []*Link{}
	}
	p.Links = append(p.Links, &Link{Type:depType, Source:p.Name, Target:target})
}

func (p *Node) findLinksByType(depType string) []*Link {
	links := []*Link{}
	for _, link := range p.Links {
		if link.Type == depType {
			links = append(links, link)
		}
	}
	return links
}

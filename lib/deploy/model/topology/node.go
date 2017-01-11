package topology

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)


//Node is element of topology
type Node struct {
	gorm.Model
	Status
	TopologyID int        `gorm:"index"` //Foreign key of the topology
	Name       string                    //node name
	Template   string                    //template name
	Class      string                    //the name of resource class
	Address    string                    //address of node.
	Credential string                    //credential key
	Links      []*Link                   //dependencies of node
	Properties []*Property               //the configuration of a node
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

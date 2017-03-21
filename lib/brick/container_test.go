package brick

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	container := NewContainer()
	container.Add(&A{}, "a", nil)
	container.Add(&B{}, "b", nil)
	container.Add(&C{}, "c", nil)
	container.Build()
	defer container.Dispose()
	obj := container.GetByName("a")
	ai, _ := obj.(AI)
	if "Bfbi event" != ai.Fai() {
		t.Error("failed")
	}
}

type AI interface {
	Fai() string
}
type A struct {
	B    *B `inject:"b"`
	C    CI `inject:"c"`
	data string
}

func (p *A) SetContainer(c *Container) {
	fmt.Printf("A.SetContainer: %+v\n", c)
}

func (p *A) Init() error {
	fmt.Println("Init")
	p.B.On("CallFbi", p)
	return nil
}

func (p *A) Handle(event string, data interface{}) {
	if s, ok := data.(string); ok {
		p.data = s
	}
}

func (p *A) Fai() string {
	fmt.Println("a.Fai")
	return p.B.Fbi() + p.data
}

type B struct {
	Trigger
}

func (p *B) SetContainer(c *Container) {
	fmt.Printf("B.SetContainer: %+v", c)
}
func (p *B) Fbi() string {
	fmt.Println("b.Fbi")
	p.Emmit("CallFbi", "fbi event")

	return "B"
}
func (p *B) FAi() string {
	return "BA"
}

type CI interface {
	Fci() string
}
type C struct {
}

func (p *C) Fci() string {
	return "C"
}

package service

import (
	"testing"
	"fmt"
)

func TestContainer(t *testing.T) {
	container := NewContainer()
	container.Add(&A{}, "a", nil)
	container.Add(&B{}, "b", nil)
	container.Add(&C{}, "c", nil)
	container.Build()
	defer container.Dispose()
	obj := container.GetByName("a")
	ai, _ := obj.(AI)
	if "B" != ai.Fai() {
		t.Error("failed")
	}
}

type AI interface {
	Fai() string
}
type A struct {
	B BI `inject:"b"`
	C CI `inject:"c"`
}

func (p *A) SetContainer(c *Container) {
	fmt.Printf("A.SetContainer: %+v", c)
}
func (p *A) Fai() string {
	return p.B.Fbi()
}

type BI interface {
	Fbi() string
}
type B struct {

}
func (p *B) SetContainer(c *Container) {
	fmt.Printf("B.SetContainer: %+v", c)
}
func (p *B) Fbi() string {
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

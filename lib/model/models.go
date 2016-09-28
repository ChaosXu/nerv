package model


var Models = map[string]*ModelDescriptor{}

type ModelDescriptor struct {
	Type interface{}
	New  func() interface{}
	NewSlice  func() interface{}
}

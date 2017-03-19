package service

import (
	"reflect"
	"fmt"
)

type ObjectState uint32

const (
	Empty ObjectState = 0
	New = 1
	Init = 2
)

// ObjectRef store the information of a object instance
type ObjectRef struct {
	objType reflect.Type
	key     string
	factory Factory
	state   ObjectState
	obj     interface{}
}

func key(objType reflect.Type, name string) string {
	key := name
	if name == "" {
		key = objType.Name()
	}
	return key
}

func newObjectRef(objType reflect.Type, name string, factory Factory) *ObjectRef {

	return &ObjectRef{objType:objType, key:key(objType, name), factory:factory, state:Empty}
}

func (p *ObjectRef) Key() string {
	return p.key
}

func (p *ObjectRef) Target() interface{} {
	return p.obj
}

func (p *ObjectRef) Type() reflect.Type {
	return p.objType
}

func (p *ObjectRef) new(obj interface{}) {
	p.obj = obj
	p.state = New
}

func (p *ObjectRef) init(obj interface{}) {
	p.obj = obj
	p.state = Init
}

// Container manage all services
type Container struct {
	objs map[string]*ObjectRef
}

func NewContainer() *Container {
	return &Container{objs:map[string]*ObjectRef{}}
}

func (p *Container) Add(obj interface{}, name string, factory Factory) {
	objType := reflect.TypeOf(obj)
	if objType.Kind()==reflect.Ptr {
		st := newObjectRef(objType, name, factory)
		p.objs[st.Key()] = st
	}else{
		panic("obj must be a pointer")
	}

}

func (p *Container) GetByName(svcType reflect.Type, name string) interface{} {
	key := key(svcType, name)
	ref := p.objs[key]
	if ref == nil {
		return nil
	}

	if ref.state == Init {
		return ref.Target()
	}

	p.initObject(ref)

	return ref.Target()
}

func (p *Container) GetByType(svcType reflect.Type) interface{} {
	return p.GetByName(svcType, "")
}

func (p *Container) initObject(r *ObjectRef) {
	factory := r.factory
	var obj interface{}
	if factory != nil {
		obj = factory.New()
	} else {
		obj = reflect.New(r.Type().Elem()).Interface()
	}

	r.new(obj)
	p.inject(r)
	r.init(obj)
}

func (p *Container) inject(r *ObjectRef) {
	t := r.Type().Elem()
	count := t.NumField()
	for i := 0; i < count; i++ {
		f := t.Field(i)

		injectObjName := f.Tag.Get("inject")

		if injectObjName != "" {
			injectR := p.objs[injectObjName]
			if injectR == nil {
				panic(fmt.Errorf("could not found object %s,defines in %s.%s", injectObjName, r.Type().Name(), f.Name))
			}
			switch injectR.state {
			case New:
				panic(fmt.Errorf("cycle dependency %s,defines in %s.%s", injectObjName, r.Type().Name(), f.Name))
			case Empty:
				p.initObject(injectR)
			}

			v := reflect.ValueOf(r.obj)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			fv := v.Field(i)
			fv.Set(reflect.ValueOf(injectR.Target()))
		}
	}
}


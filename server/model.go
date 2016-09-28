package main

import (
	"github.com/chaosxu/nerv/lib/model"
)


func AutoMigrate() {
	for _, v := range model.Models {
		DB.AutoMigrate(v.Type)
	}
}


//func ServiceTemplate() *ModelDescriptor {
//	return &ModelDescriptor{
//		Type: &model.ServiceTemplate{},
//		New: func() interface{} {
//			return &model.ServiceTemplate{}
//		},
//	}
//}
//
//func Topology() *ModelDescriptor {
//	return &ModelDescriptor{
//		Type: &model.Topology{},
//		New: func() interface{} {
//			return &model.Topology{}
//		},
//	}
//}
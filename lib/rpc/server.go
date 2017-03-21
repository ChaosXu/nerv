package rpc

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"reflect"
)

var (
	receivers []interface{} = []interface{}{} //RPC service registry
)

//Register a rpc service
func Register(service interface{}) {
	receivers = append(receivers, service)
}

//Start rpc server and register all rpc handlers from var Receivers.
func Start(cfg *env.Properties) error {
	port := cfg.GetMapString("rpc", "port", "3334")
	if port == "" {
		return fmt.Errorf("rpc_port isn't setted")
	}

	for _, rcvr := range receivers {
		if err := rpc.Register(rcvr); err != nil {
			return err
		} else {
			log.Printf("Register %s\n", reflect.TypeOf(rcvr).String())
		}
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		return nil
	}

	return http.Serve(l, nil)
}

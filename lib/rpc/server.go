package rpc

import (
	"github.com/ChaosXu/nerv/lib/env"
	"net"
	"net/rpc"
	"log"
	"fmt"
	"reflect"
	"net/http"
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
	l, e := net.Listen("tcp", ":" + port)
	if e != nil {
		return nil
	}

	go func() {
		if err := http.Serve(l, nil); err != nil {
			fmt.Println(err.Error())
		}
	}()
	return nil
}

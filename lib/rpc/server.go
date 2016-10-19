package rpc

import (
	"github.com/ChaosXu/nerv/lib/env"
	"net"
	"net/rpc"
	"log"
	"net/rpc/jsonrpc"
	"fmt"
)

var (
	receivers []interface{} = []interface{}{} //RPC service registry
)

//Register a rpc service
func Register(service interface{}) {
	receivers = append(receivers, service)
}

//Start rpc server and register all rpc handlers from var Receivers.
func Start() error {
	port := env.Config().GetString("rpc_port")
	if port == "" {
		return fmt.Errorf("rpc_port isn't setted")
	}

	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return err
	}
	defer listener.Close()

	srv := rpc.NewServer()
	for rcvr := range receivers {
		if err := srv.Register(rcvr); err != nil {
			return err
		} else {
			log.Printf("Register %+v\n", rcvr)
		}
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error. %s\n", err.Error())
			continue
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

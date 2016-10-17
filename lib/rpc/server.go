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
	Receivers map[string]interface{} = map[string]interface{}{}
)

//Start rpc server and register all rpc handlers from var Receivers.
func Start() error {
	port := env.Config.GetProperty("rpc_port")
	if port == "" {
		return fmt.Errorf("rpc_port isn't setted")
	}

	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return err
	}
	defer listener.Close()

	srv := rpc.NewServer()
	for name, rcvr := range Receivers {
		if err := srv.RegisterName(name, rcvr); err != nil {
			return err
		} else {
			log.Printf("Register %s\n", name)
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

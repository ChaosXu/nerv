package service

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/rpc"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// RemoteScriptServiceFactory
type RemoteScriptServiceFactory struct {
}

func (p *RemoteScriptServiceFactory) New() interface{} {
	agent, err := NewRemoteScriptService(env.Config())
	if err != nil {
		panic(err)
	}
	return agent
}

//RemoteScriptService execute the method of app
type Agent struct {
	AppRoot string //the root path of app
	cfg     *env.Properties
}

func NewRemoteScriptService(cfg *env.Properties) (*Agent, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	appRoot := cfg.GetMapString("app", "root", "../app")
	appRoot = filepath.Join(dir, appRoot)
	if err := os.MkdirAll(appRoot, os.ModeDir|os.ModePerm); err != nil {
		return nil, err
	}

	return &Agent{AppRoot: appRoot, cfg: cfg}, nil
}

func (p *Agent) Init() error {
	rpc.Register(p)
	go func() {
		log.Fatal(rpc.Start(p.cfg))
	}()
	return nil
}

//Execute a script in the host of the agent
func (p *Agent) Execute(script *rpc.RemoteScript, reply *string) error {
	fmt.Println("Agent.Execute")
	//Optimize: async

	fmt.Println(script.Content)

	out, err := exec.Command("/bin/bash", "-c", script.Content).Output()
	if err != nil {
		res := string(out)
		fmt.Println("err:" + res)
		return err
	}
	res := string(out)
	fmt.Println(res)
	*reply = res
	return nil
}

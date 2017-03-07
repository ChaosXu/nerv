package service

import (
	"os"
	"log"
	"path/filepath"
	"github.com/ChaosXu/nerv/lib/env"
	"fmt"
	"os/exec"
	"github.com/ChaosXu/nerv/lib/rpc"
	"github.com/ChaosXu/nerv/lib/service"
)

func init() {
	service.Registry.Put("RemoteScript", &RemoteScriptServiceFactory{})
}

type RemoteScriptServiceFactory struct {
	remoteScriptService *Agent
}

func (p *RemoteScriptServiceFactory) Init() error {
	agent, err := NewRemoteScriptService(env.Config())
	if err != nil {
		return err
	}
	p.remoteScriptService = agent
	return nil
}

func (p *RemoteScriptServiceFactory) Get() interface{} {
	return p.remoteScriptService
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
	if err := os.MkdirAll(appRoot, os.ModeDir | os.ModePerm); err != nil {
		return nil, err
	}

	return &Agent{AppRoot:appRoot, cfg:cfg}, nil
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




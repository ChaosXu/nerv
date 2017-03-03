package service

import (
	"os"
	"log"
	"path/filepath"
	"github.com/ChaosXu/nerv/lib/env"
	"fmt"
	"os/exec"
	"github.com/ChaosXu/nerv/lib/rpc"
)

//RemoteScriptService execute the method of app
type RemoteScriptService struct {
	AppRoot string //the root path of app
	cfg     *env.Properties
}

func NewRemoteScriptService(cfg *env.Properties) (*RemoteScriptService, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	appRoot := cfg.GetMapString("app", "root", "../app")
	appRoot = filepath.Join(dir, appRoot)
	if err := os.MkdirAll(appRoot, os.ModeDir | os.ModePerm); err != nil {
		return nil, err
	}

	return &RemoteScriptService{AppRoot:appRoot, cfg:cfg}, nil
}

func (p *RemoteScriptService) Start() error {
	rpc.Register(p)
	return rpc.Start(p.cfg)
}

//Execute a script in the host of the agent
func (p *RemoteScriptService) Execute(script *rpc.RemoteScript, reply *string) error {
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




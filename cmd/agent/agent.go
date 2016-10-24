package main

import (
	"os"
	"log"
	"path/filepath"
	"github.com/ChaosXu/nerv/lib/env"
	"fmt"
	"os/exec"
)

//Agent execute the method of app
type Agent struct {
	AppRoot string //the root path of app
}

func NewAgent() (*Agent, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	appRoot := env.Config().GetMapString("app", "root", "../app")
	appRoot = filepath.Join(dir, appRoot)
	if err := os.MkdirAll(appRoot, os.ModeDir | os.ModePerm); err != nil {
		return nil, err
	}

	return &Agent{AppRoot:appRoot}, nil
}

type RemoteScript struct {
	Content string
	Args    map[string]string
}

//Execute a script in the host of the agent
func (p *Agent) Execute(script *RemoteScript, response *string) error {
	export := ""
	for k, v := range script.Args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	export = export + " APP_ROOT=" + p.AppRoot

	shell := "export " + export + " && " + script.Content
	log.Println(shell)

	out, err := exec.Command("/bin/bash", "-c", shell).Output()
	if err != nil {
		return err
	}
	log.Println(string(out))

	return nil
}




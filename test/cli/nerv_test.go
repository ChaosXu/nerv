package cli

import (
	"testing"
	"regexp"
	"net/rpc"
)

type RemoteScript struct {
	Content string
	Args    map[string]string
}

func TestNervCmd(t *testing.T) {

	//create
	cmd := &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "create", "-t", "../../resources/templates/nerv/env_standalone.json", "-o", "nerv-test", "-n", "../../../../test/cli/nerv_standalone_inputs.json"},
	}

	var id string
	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		res := string(out)
		t.Log(res)
		regex := regexp.MustCompile(`.*id=(.+)`)
		id = regex.FindStringSubmatch(res)[1]
	}

	//install
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "install", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		res := string(out)
		t.Log(res)
	}


	//setup
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "setup", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//start
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "start", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//agent test
	testAgent(t);

	//stop
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "stop", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//uninstall
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "uninstall", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}


	//delete
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "delete", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}

func testAgent(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:3334")
	if err != nil {
		t.Log("DialHTTP:", err.Error())
	}

	script := &RemoteScript{Content:"echo agnet ok", Args:map[string]string{}}
	var reply string
	err = client.Call("Agent.Execute", script, &reply)
	if err != nil {
		t.Log("Agent.Execute failed.", err)
	}else{
		t.Log(reply)
	}
}

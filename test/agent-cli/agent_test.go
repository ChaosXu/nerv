package agent_cli

import (
	"testing"
	"regexp"
	"net/rpc"
	"github.com/ChaosXu/nerv/test/util"
)

func TestNervCmd(t *testing.T) {

	//create
	cmd := &util.Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "create", "-t", "../../resources/templates/nerv/server_core.json", "-o", "nerv-test", "-n", "../../../../test/agent-cli/nerv_standalone_inputs.json"},
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
	cmd = &util.Cmd{
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
	cmd = &util.Cmd{
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
	cmd = &util.Cmd{
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

	testAgent(t);

	//stop
	cmd = &util.Cmd{
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
	//cmd = &util.Cmd{
	//	Dir: "../../release/nerv/nerv-cli/bin",
	//	Cli:"./nerv-cli",
	//	Items:[]string{"nerv", "uninstall", "-i", id},
	//}
	//
	//if out, err := cmd.Run(t); err != nil {
	//	t.Log(string(out))
	//	t.Errorf("%s", err.Error())
	//} else {
	//	t.Log(string(out))
	//}


	//delete
	//cmd = &util.Cmd{
	//	Dir: "../../release/nerv/nerv-cli/bin",
	//	Cli:"./nerv-cli",
	//	Items:[]string{"nerv", "delete", "-i", id},
	//}
	//
	//if out, err := cmd.Run(t); err != nil {
	//	t.Log(string(out))
	//	t.Errorf("%s", err.Error())
	//} else {
	//	t.Log(string(out))
	//}
}

func testAgent(t *testing.T) {
	testHttp(t)
	testAppCmd(t)
}

func testHttp(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:3334")
	if err != nil {
		t.Log("DialHTTP:", err.Error())
	}

	script := &util.RemoteScript{Content:"echo agnet ok", Args:map[string]string{}}
	var reply string
	err = client.Call("Agent.Execute", script, &reply)
	if err != nil {
		t.Log("Agent.Execute failed.", err)
	} else {
		t.Log(reply)
	}
}

func testAppCmd(t *testing.T) {
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "create", "-D", "../../../../test/agent-cli/app.json"})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "list"})
}

func runCmd(t *testing.T, dir string, cli string, args []string) {
	cmd := &util.Cmd{
		Dir: dir,
		Cli:cli,
		Items:args,
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}


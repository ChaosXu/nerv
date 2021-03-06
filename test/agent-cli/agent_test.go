package agent_cli

import (
	"github.com/ChaosXu/nerv/test/util"
	"net/rpc"
	"regexp"
	"testing"
)

func TestNervCmd(t *testing.T) {

	//create
	cmd := &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "create", "-t", "../../resources/templates/nerv/server_core.json", "-o", "nerv-test", "-n", "../../../../test/agent-cli/nerv_standalone_inputs.json"},
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
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "install", "-i", id},
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
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "setup", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//start
	cmd = &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "start", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	testAgent(t)

	//stop
	cmd = &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "stop", "-i", id},
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

	script := &util.RemoteScript{Content: "echo agnet ok", Args: map[string]string{}}
	var reply string
	err = client.Call("Agent.Execute", script, &reply)
	if err != nil {
		t.Log("Agent.Execute failed.", err)
	} else {
		t.Log(reply)
	}
}

func testAppCmd(t *testing.T) {
	id := runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "create", "-D", "../../../../test/agent-cli/app.json"})
	t.Log(id)
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "list"})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "get", "-i", id})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "update", "-n test_app", `-a [\"path\"]`, `-v [\"a\"]`})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "get", "-i", id})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "delete", "-i", id})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "list"})
	runCmd(t, "../../release/nerv/agent/bin", "./agent-cli", []string{"app", "get", "-i", id})
}

func runCmd(t *testing.T, dir string, cli string, args []string) string {
	r := "0"
	cmd := &util.Cmd{
		Dir:   dir,
		Cli:   cli,
		Items: args,
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {

		res := string(out)
		regex := regexp.MustCompile(`.*([0-9]+),.*`)
		match := regex.FindStringSubmatch(res)
		if len(match) > 0 {
			r = match[1]
			t.Log(r)
		}
		t.Log(res)
	}
	return r
}

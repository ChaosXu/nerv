package cli

import (
	"github.com/ChaosXu/nerv/test/util"
	"regexp"
	"testing"
)

func TestNerInstall(t *testing.T) {

	//create
	cmd := &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"nerv", "create", "-t", "../../resources/templates/nerv/server_core.json", "-o", "nerv-test", "-n", "../../../../test/cli/nerv_standalone_inputs.json"},
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
}

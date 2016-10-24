package shell_test

import (
	"testing"
	"github.com/ChaosXu/nerv/lib/deploy/driver/ssh"
	"github.com/ChaosXu/nerv/lib/env"
)

func TestExecute(t *testing.T) {
	if c, err := env.LoadConfig("ssh_test_config.json"); err != nil {
		t.Fatal(err.Error())
	} else {
		env.SetConfig(c)
	}

	args := map[string]string{
		"PKG_URL":"http://10.211.55.2:3332/pkg/agent.tar.gz",
		"ROOT":"/opt",
	}
	err := ssh.Execute(
		"centos7.chaosxu.com:22",
		"_", "_",
		"http://10.211.55.2:3332/scripts/nerv/Agent/create.sh",
		args)
	if err != nil {
		t.Fatal(err.Error())
	}
}

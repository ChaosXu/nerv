package environment

import (
	"os/exec"
	"fmt"
	"github.com/ChaosXu/nerv/lib/resource/model"

	"github.com/ChaosXu/nerv/lib/resource/repository"
)

// StandaloneEnvironment where nerv deployed in standalone mode
type StandaloneEnvironment struct {
	ScriptRepository repository.ScriptRepository `inject:"script_rep_standalone"`
}

func (p *StandaloneEnvironment) Exec(class *model.Class, operation *model.Operation, args map[string]string) error {
	fmt.Println("Standalone.Exec " + class.Name + "." + operation.Name)

	script, err := p.ScriptRepository.Get(operation.Implementor)
	if err != nil {
		return err;
	}
	return p.call(script, args)
}

func (p *StandaloneEnvironment) call(script *model.Script, args map[string]string) error {
	export := ""
	for k, v := range args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}

	shell := "export " + export + " && " + script.Content
	fmt.Println(shell)

	out, err := exec.Command("/bin/bash", "-c", shell).Output()
	if err != nil {
		res := string(out)
		fmt.Println(res)
		return err
	}
	res := string(out)
	fmt.Println(res)

	return nil
}

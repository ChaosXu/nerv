package lib

import (
	"github.com/ChaosXu/nerv/lib/automation/manager"
	"github.com/facebookgo/inject"
	"github.com/ChaosXu/nerv/lib/resource/environment"
	"github.com/ChaosXu/nerv/lib/automation/repository"
	"github.com/ChaosXu/nerv/lib/db"
	resrep "github.com/ChaosXu/nerv/lib/resource/repository"
)

func NewDeployer() (*manager.Deployer, error) {
	var g inject.Graph
	var deployer manager.Deployer
	var templateRep repository.LocalTemplateRepository
	var dbService db.DBService
	var executor environment.ExecutorImpl
	classRep := resrep.NewStandaloneClassRepository("../../resources/scripts")
	scriptRep := resrep.NewStandaloneScriptRepository("../../resources/scripts")
	standaloneEnv := environment.StandaloneEnvironment{ScriptRepository:scriptRep}
	sshEnv := environment.SshEnvironment{ScriptRepository:scriptRep}
	err := g.Provide(
		&inject.Object{Value: &deployer},
		&inject.Object{Value: &templateRep},
		&inject.Object{Value: &dbService},
		&inject.Object{Value: &executor},
		&inject.Object{Value: &standaloneEnv, Name:"env_standalone"},
		&inject.Object{Value: &sshEnv, Name:"env_ssh"},
		&inject.Object{Value: classRep},
	)
	if err != nil {
		return nil, err
	}

	err = g.Populate()
	if err != nil {
		return nil, err
	}
	return &deployer, nil
}

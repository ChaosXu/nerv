package lib

import (
	"github.com/ChaosXu/nerv/lib/automation/manager"
	"github.com/facebookgo/inject"
	"github.com/ChaosXu/nerv/lib/operation"
	"github.com/ChaosXu/nerv/lib/automation/repository"
	"github.com/ChaosXu/nerv/lib/db"
	resrep "github.com/ChaosXu/nerv/lib/resource/repository"
	"github.com/ChaosXu/nerv/lib/service"
	"fmt"
)

func init() {
	service.Registry.Put("Topology", &TopologyServiceFactory{})
}

type TopologyServiceFactory struct {
	deployer *manager.Deployer
}

func (p *TopologyServiceFactory) Init() error {
	deployer, err := p.newDeployer()
	if err != nil {
		return err
	}
	p.deployer = deployer
	return nil
}

func (p *TopologyServiceFactory) Get() (interface{}, error) {
	if p.deployer == nil {
		return nil, fmt.Errorf("%s is uninit", "deployService")
	}
	return p.deployer, nil
}

func (p *TopologyServiceFactory) newDeployer() (*manager.Deployer, error) {
	var g inject.Graph
	var deployer manager.Deployer
	var templateRep repository.HttpTemplateRepository
	var dbService db.DBService
	var executor operation.ExecutorImpl
	classRep := resrep.NewHttpClassRepository("../../resources/scripts")
	scriptRep := resrep.NewHttpScriptRepository("../../resources/scripts")
	standaloneEnv := operation.StandaloneEnvironment{ScriptRepository:scriptRep}
	sshEnv := operation.SshEnvironment{ScriptRepository:scriptRep}
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

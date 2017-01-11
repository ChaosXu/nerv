package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/facebookgo/inject"
	"github.com/ChaosXu/nerv/lib/deploy/repository"
	"github.com/ChaosXu/nerv/lib/deploy/manager"
	classrep "github.com/ChaosXu/nerv/lib/resource/repository"
	"github.com/ChaosXu/nerv/lib/resource/executor"
)

var init_flag_template string
var init_flag_topology_name string
var init_flag_config string

func init() {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Install nerv in one host",
		Long:
		`
		Use the init command to install all components of nerv in one host
		`,
		RunE: install,
	}
	initCmd.Flags().StringVarP(&init_flag_template, "template", "t", "", "required. The path of template that used to install nerv")
	initCmd.Flags().StringVarP(&init_flag_topology_name, "topolgoy", "o", "", "required. Topology name")
	initCmd.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")

	RootCmd.AddCommand(initCmd)
}

func install(cmd *cobra.Command, args []string) error {
	if init_flag_template == "" {
		return errors.New("--template -t is null")
	}

	if init_flag_topology_name == "" {
		return errors.New("--topology -o is null")
	}

	//init
	env.InitByConfig(init_flag_config)
	initDB()
	defer db.DB.Close()

	var g inject.Graph
	var manager manager.Deployer
	var templateRep repository.LocalTemplateRepository
	var dbService db.DBService
	classRep := classrep.NewLocalClassRepository("../../resources/scripts")
	var Executor executor.ExecutorImpl
	err := g.Provide(
		&inject.Object{Value: &manager},
		&inject.Object{Value: &templateRep},
		&inject.Object{Value: &dbService},
		&inject.Object{Value: classRep},
		&inject.Object{Value: &Executor},
	)
	if err != nil {
		return err
	}

	err = g.Populate()
	if err != nil {
		return err
	}
	return manager.Install(init_flag_topology_name, init_flag_template)
}

func initDB() {
	url := fmt.Sprintf(
		"%s:%s@%s",
		env.Config().GetMapString("db", "user", "root"),
		env.Config().GetMapString("db", "password", "root"),
		env.Config().GetMapString("db", "url"),
	)
	gdb, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db.DB = gdb
	db.DB.LogMode(true)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
}



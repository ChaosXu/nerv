package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/cmd/cli/lib"
)

var init_flag_template string
var init_flag_topology_name string
var init_flag_config string

func init() {
	var topo = &cobra.Command{
		Use:    "topo [command] [flags]",
		Short:    "Manage the topology resource",
		Long:    "Manage the topology resource",
		RunE: topo,
	}
	RootCmd.AddCommand(topo)

	var create = &cobra.Command{
		Use:    "create",
		Short:    "Create a topology",
		Long:    "Create a topology",
		RunE: create,
	}
	create.Flags().StringVarP(&init_flag_template, "template", "t", "", "required. The path of template that used to install nerv")
	create.Flags().StringVarP(&init_flag_topology_name, "topolgoy", "o", "", "required. Topology name")
	create.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(create)

}

func topo(cmd *cobra.Command, args []string) error {
	return nil
}

func create(cmd *cobra.Command, args []string) error {
	if init_flag_template == "" {
		return errors.New("--template -t is null")
	}

	if init_flag_topology_name == "" {
		return errors.New("--topology -o is null")
	}

	//init
	env.InitByConfig(init_flag_config)
	db := lib.InitDB()
	defer db.Close()

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}
	return deployer.Install(init_flag_topology_name, init_flag_template)
}





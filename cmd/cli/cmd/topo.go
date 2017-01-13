package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/cmd/cli/lib"
	"fmt"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
)

var init_flag_template string
var init_flag_topology_name string
var init_flag_config string
var init_flag_id uint

func init() {
	var topo = &cobra.Command{
		Use:    "topo [command] [flags]",
		Short:    "Manage the topology resource",
		Long:    "Manage the topology resource",
		RunE: topo,
	}
	RootCmd.AddCommand(topo)

	//list
	var list = &cobra.Command{
		Use:    "list",
		Short:    "List all topologies",
		Long:    "List all topologies",
		RunE: list,
	}
	list.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(list)


	//create
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


	//delete
	var delete = &cobra.Command{
		Use:    "delete",
		Short:    "Delete a topology",
		Long:    "Delete a topology",
		RunE: remove,
	}
	delete.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	delete.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(delete)

	//install
	var install = &cobra.Command{
		Use:    "install",
		Short:    "Install a topology to an environment",
		Long:    "Install a topology to an environment",
		RunE: install,
	}
	install.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	install.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(install)

}

func topo(cmd *cobra.Command, args []string) error {
	return nil
}

// create a topology
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
	id, err := deployer.Create(init_flag_topology_name, init_flag_template)
	if err != nil {
		return err;
	}

	fmt.Printf("Topology has been created. id=%d\n", id)
	return nil
}

// list all topologies
func list(cmd *cobra.Command, args []string) error {
	//init
	env.InitByConfig(init_flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	topos := []topology.Topology{}
	if err := gdb.Find(&topos).Error; err != nil {
		return err
	}

	fmt.Println("ID\tName\tRunStatus\tCreateAt\tTemplate")
	for _, topo := range topos {
		fmt.Printf("%d\t%s\t%d\t%s\t%s\n", topo.ID, topo.Name, topo.RunStatus, topo.CreatedAt.Format("2006-01-02 15:04:05"), topo.Template)
	}
	return nil
}

// delete a topology
func remove(cmd *cobra.Command, args []string) error {
	if init_flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(init_flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := init_flag_id
	data := topology.Topology{}
	if err := gdb.First(&data, id).Error; err != nil {
		return err
	}

	if err := gdb.Unscoped().Delete(data).Error; err != nil {
		return err
	}

	fmt.Printf("Topology has been deleted. id=%d\n", id)

	return nil
}

// Install a topology
func install(cmd *cobra.Command, args []string) error {
	if init_flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(init_flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := init_flag_id
	topo := &topology.Topology{}
	if err := gdb.First(topo, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Install(topo.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Topology has been installed. id=%d\n", id)

	return nil
}





package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/cmd/cli/lib"
	"fmt"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
	"encoding/json"
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

	//get
	var get = &cobra.Command{
		Use:    "get",
		Short:    "Get a topology",
		Long:    "Get all topology",
		RunE: get,
	}
	get.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	get.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(get)


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

	//uninstall
	var uninstall = &cobra.Command{
		Use:    "uninstall",
		Short:    "Uninstall a topology from an environment",
		Long:    "Uninstall a topology from an environment",
		RunE: uninstall,
	}
	uninstall.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	uninstall.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(uninstall)

	//start
	var start = &cobra.Command{
		Use:    "start",
		Short:    "Start a topology from an environment",
		Long:    "Start a topology from an environment",
		RunE: start,
	}
	start.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	start.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(start)

	//stop
	var stop = &cobra.Command{
		Use:    "stop",
		Short:    "Stop a topology from an environment",
		Long:    "Stop a topology from an environment",
		RunE: stop,
	}
	stop.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	stop.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(stop)

	//restart
	var restart = &cobra.Command{
		Use:    "restart",
		Short:    "Restart a topology from an environment",
		Long:    "Restart a topology from an environment",
		RunE: restart,
	}
	restart.Flags().UintVarP(&init_flag_id, "id", "i", 0, "Topology id")
	restart.Flags().StringVarP(&init_flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(restart)
}

func topo(cmd *cobra.Command, args []string) error {
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

// get a topologies
func get(cmd *cobra.Command, args []string) error {
	//init
	env.InitByConfig(init_flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := init_flag_id
	data := topology.Topology{}
	if err := gdb.Preload("Nodes").Preload("Nodes.Links").Preload("Nodes.Properties").First(&data, id).Error; err != nil {
		return err
	}
	buf, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(buf))
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

	fmt.Printf("Create topology success. id=%d\n", id)
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

	fmt.Printf("Delete topology success. id=%d\n", id)

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

	fmt.Printf("Install topology success. id=%d\n", id)

	return nil
}

// Uninstall a topology
func uninstall(cmd *cobra.Command, args []string) error {
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

	err = deployer.Uninstall(topo.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Uninstall topology success. id=%d\n", id)

	return nil
}

// Start a topology
func start(cmd *cobra.Command, args []string) error {
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

	err = deployer.Start(topo.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Start topology success. id=%d\n", id)

	return nil
}

// Stop a topology
func stop(cmd *cobra.Command, args []string) error {
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

	err = deployer.Stop(topo.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Stop topology success. id=%d\n", id)

	return nil
}

// Restart a topology
func restart(cmd *cobra.Command, args []string) error {
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

	err = deployer.Restart(topo.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Restart topology success. id=%d\n", id)

	return nil
}





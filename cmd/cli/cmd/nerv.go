package cmd

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/toolkits/file"

	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/cmd/cli/lib"
	"github.com/ChaosXu/nerv/lib/automation/model/topology"
	"github.com/ChaosXu/nerv/lib/cli"
)

var NervCmd = &cobra.Command{
	Use:    "nerv [command] [flags]",
	Short:    "Manage the platform",
	Long:    "Manage the platform",
	RunE: nervCmd,
}

func init() {

	RootCmd.AddCommand(NervCmd)

	//list
	var list = &cobra.Command{
		Use:    "list",
		Short:    "List all platforms",
		Long:    "List all platforms",
		RunE: listNerv,
	}
	list.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(list)

	//get
	var get = &cobra.Command{
		Use:    "get",
		Short:    "Get a platform",
		Long:    "Get all platform",
		RunE: getNerv,
	}
	get.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	get.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(get)


	//create
	var create = &cobra.Command{
		Use:    "create",
		Short:    "Create a platform",
		Long:    "Create a platform",
		RunE: createNerv,
	}
	create.Flags().StringVarP(&cli.Flag_template, "template", "t", "../../resources/templates/nerv/env_standalone.json", "required. The path of template that used to install nerv")
	create.Flags().StringVarP(&cli.Flag_topology_name, "topologoy", "o", "nerv-standalone", "required. Topology name")
	create.Flags().StringVarP(&cli.Flag_input_path, "input", "n", "", "required. The path of input that a template need it as input arguments")
	create.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(create)


	//delete
	var delete = &cobra.Command{
		Use:    "delete",
		Short:    "Delete a platform",
		Long:    "Delete a platform",
		RunE: removeNerv,
	}
	delete.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	delete.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(delete)

	//install
	var install = &cobra.Command{
		Use:    "install",
		Short:    "Install a platform to an environment",
		Long:    "Install a platform to an environment",
		RunE: installNerv,
	}
	install.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	install.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(install)

	//uninstall
	var uninstall = &cobra.Command{
		Use:    "uninstall",
		Short:    "Uninstall a platform from an environment",
		Long:    "Uninstall a platform from an environment",
		RunE: uninstallNerv,
	}
	uninstall.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	uninstall.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(uninstall)

	//start
	var start = &cobra.Command{
		Use:    "start",
		Short:    "Start a platform from an environment",
		Long:    "Start a platform from an environment",
		RunE: startNerv,
	}
	start.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	start.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(start)

	//stop
	var stop = &cobra.Command{
		Use:    "stop",
		Short:    "Stop a platform from an environment",
		Long:    "Stop a platform from an environment",
		RunE: stopNerv,
	}
	stop.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	stop.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(stop)

	//restart
	var restart = &cobra.Command{
		Use:    "restart",
		Short:    "Restart a platform from an environment",
		Long:    "Restart a platform from an environment",
		RunE: restartNerv,
	}
	restart.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	restart.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(restart)

	//setup
	var setup = &cobra.Command{
		Use:    "setup",
		Short:    "Setup configuration",
		Long:    "Setup configuration of all nodes in platform",
		RunE: setupNerv,
	}
	setup.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	setup.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	NervCmd.AddCommand(setup)
}

func nervCmd(cmd *cobra.Command, args []string) error {
	return nil
}

func listNerv(cmd *cobra.Command, args []string) error {
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	nervs := []topology.Topology{}
	if err := gdb.Find(&nervs).Error; err != nil {
		return err
	}

	fmt.Println("ID\tName\tVersion\tRunStatus\tCreateAt\tTemplate")
	for _, nerv := range nervs {
		fmt.Printf("%d\t%s\t%d\t%d\t%s\t%s\n", nerv.ID, nerv.Name, nerv.Version, nerv.RunStatus, nerv.CreatedAt.Format("2006-01-02 15:04:05"), nerv.Template)
	}
	return nil
}

func getNerv(cmd *cobra.Command, args []string) error {
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
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

func createNerv(cmd *cobra.Command, args []string) error {
	var inputs map[string]interface{}
	if cli.Flag_input_path != "" {
		buf, err := file.ToBytes(cli.Flag_input_path)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(buf, &inputs); err != nil {
			return err
		}
	}
	//init
	env.InitByConfig(cli.Flag_config)
	db := lib.InitDB()
	defer db.Close()

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}
	id, err := deployer.Create(cli.Flag_topology_name, cli.Flag_template, inputs)
	if err != nil {
		return err;
	}

	fmt.Printf("Create platform success. id=%d\n", id)
	return nil
}

func removeNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	data := topology.Topology{}
	if err := gdb.First(&data, id).Error; err != nil {
		return err
	}

	if err := gdb.Unscoped().Delete(data).Error; err != nil {
		return err
	}

	fmt.Printf("Delete platform success. id=%d\n", id)

	return nil
}

func installNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Install(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Install platform success. id=%d\n", id)

	return nil
}

func uninstallNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Uninstall(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Uninstall platform success. id=%d\n", id)

	return nil
}

func startNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Start(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Start platform success. id=%d\n", id)

	return nil
}

func stopNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Stop(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Stop platform success. id=%d\n", id)

	return nil
}

func restartNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Restart(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Restart platform success. id=%d\n", id)

	return nil
}

func setupNerv(cmd *cobra.Command, args []string) error {
	if cli.Flag_id == 0 {
		return errors.New("--id -i is null")
	}
	//init
	env.InitByConfig(cli.Flag_config)
	gdb := lib.InitDB()
	defer gdb.Close()

	id := cli.Flag_id
	nerv := &topology.Topology{}
	if err := gdb.First(nerv, id).Error; err != nil {
		return err
	}

	deployer, err := lib.NewDeployer()
	if err != nil {
		return err
	}

	err = deployer.Setup(nerv.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Setup platform success. id=%d\n", id)

	return nil
}





package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/cli/format"
	"github.com/ChaosXu/nerv/lib/cli"
)

func init() {
	var topo = &cobra.Command{
		Use:    "topo [command] [flags]",
		Aliases: []string{"Topology"},
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
		RunE: cli.ListObjsFunc("Topology",
			&format.Page{List:"data", Columns:[]format.Column{
				{Name:"ID", Format:"%v"},
				{Name:"name", Label:"Name", Format:"%s"},
				{Name:"version", Label:"Version", Format:"%v"},
				{Name:"RunStatus", Format:"%v"},
				{Name:"Error", Format:"%s"},
				{Name:"template", Label:"Template", Format:"%s"},
			}}),
	}
	list.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(list)

	//get
	var get = &cobra.Command{
		Use:    "get",
		Short:    "Get a topology",
		Long:    "Get all topology",
		RunE: cli.GetObjFunc("Topology", []string{"Nodes", "Nodes.Links", "Nodes.Properties"}),
	}
	get.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	get.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(get)


	//create
	var create = &cobra.Command{
		Use:    "create",
		Short:    "Create a topology",
		Long:    "Create a topology",
		RunE: cli.InvokeSvcFunc("Topology", "Create", []cli.ArgType{{Flag:"topology", Type:"string"}, {Flag:"template", Type:"string"}, {Flag:"input", Type:"ref"}}),
	}
	create.Flags().StringVarP(&cli.Flag_topology_name, "topology", "o", "", "required. Topology name")
	create.Flags().StringVarP(&cli.Flag_template, "template", "t", "", "required. The path of template that used to install nerv")
	create.Flags().StringVarP(&cli.Flag_input_path, "input", "n", "", "required. The path of input that a template need it as input arguments")
	create.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(create)

	//migrate
	var migrate = &cobra.Command{
		Use:    "migrate",
		Short:    "Migrate a topology",
		Long:    "Migreate a topology for scaling out of scaling in a service",
		RunE: cli.InvokeSvcFunc("Topology", "Migrate", []cli.ArgType{{Flag:"id", Type:"uint"}, {Flag:"input", Type:"ref"}}),
	}
	migrate.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	migrate.Flags().StringVarP(&cli.Flag_input_path, "input", "n", "", "required. The path of input that a template need it as input arguments")
	migrate.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(migrate)


	//delete
	var delete = &cobra.Command{
		Use:    "delete",
		Short:    "Delete a topology",
		Long:    "Delete a topology",
		RunE: cli.RemoveObjFunc("Topology"),
	}
	delete.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	delete.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(delete)

	//install
	var install = &cobra.Command{
		Use:    "install",
		Short:    "Install a topology to an environment",
		Long:    "Install a topology to an environment",
		RunE: cli.InvokeSvcFunc("Topology", "Install", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	install.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	install.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(install)

	//uninstall
	var uninstall = &cobra.Command{
		Use:    "uninstall",
		Short:    "Uninstall a topology from an environment",
		Long:    "Uninstall a topology from an environment",
		RunE: cli.InvokeSvcFunc("Topology", "Uninstall", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	uninstall.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	uninstall.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(uninstall)

	//reload
	var start = &cobra.Command{
		Use:    "reload",
		Short:    "Reload a topology that activate new config",
		Long:    "Reload a topology that activate new config",
		RunE: cli.InvokeSvcFunc("Topology", "Reload", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	start.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	start.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(start)

	//stop
	var stop = &cobra.Command{
		Use:    "stop",
		Short:    "Stop a topology from an environment",
		Long:    "Stop a topology from an environment",
		RunE: cli.InvokeSvcFunc("Topology", "Stop", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	stop.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	stop.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(stop)

	//start
	var restart = &cobra.Command{
		Use:    "start",
		Short:    "Start a topology from an environment",
		Long:    "Start a topology from an environment",
		RunE: cli.InvokeSvcFunc("Topology", "Start", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	restart.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	restart.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(restart)

	//setup
	var setup = &cobra.Command{
		Use:    "setup",
		Short:    "Setup configuration",
		Long:    "Setup configuration of all nodes in topology",
		RunE: cli.InvokeSvcFunc("Topology", "Setup", []cli.ArgType{{Flag:"id", Type:"uint"}}),
	}
	setup.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "Topology id")
	setup.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(setup)
}

func topo(cmd *cobra.Command, args []string) error {
	return nil
}







package cmd

import "github.com/spf13/cobra"

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
		RunE: listObjs,
	}
	list.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(list)

	//get
	var get = &cobra.Command{
		Use:    "get",
		Short:    "Get a topology",
		Long:    "Get all topology",
		RunE: getObjFunc([]string{"Nodes", "Nodes.Links", "Nodes.Properties"}),
	}
	get.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	get.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(get)


	//create
	var create = &cobra.Command{
		Use:    "create",
		Short:    "Create a topology",
		Long:    "Create a topology",
		RunE: createObj,
	}
	create.Flags().StringVarP(&flag_template, "template", "t", "", "required. The path of template that used to install nerv")
	create.Flags().StringVarP(&flag_topology_name, "topolgoy", "o", "", "required. Topology name")
	create.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(create)


	//delete
	var delete = &cobra.Command{
		Use:    "delete",
		Short:    "Delete a topology",
		Long:    "Delete a topology",
		RunE: removeObj,
	}
	delete.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	delete.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(delete)

	//install
	var install = &cobra.Command{
		Use:    "install",
		Short:    "Install a topology to an environment",
		Long:    "Install a topology to an environment",
		RunE: invokeObjFunc("install", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	install.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	install.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(install)

	//uninstall
	var uninstall = &cobra.Command{
		Use:    "uninstall",
		Short:    "Uninstall a topology from an environment",
		Long:    "Uninstall a topology from an environment",
		RunE: invokeObjFunc("uninstall", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	uninstall.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	uninstall.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(uninstall)

	//start
	var start = &cobra.Command{
		Use:    "start",
		Short:    "Start a topology from an environment",
		Long:    "Start a topology from an environment",
		RunE: invokeObjFunc("start", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	start.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	start.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(start)

	//stop
	var stop = &cobra.Command{
		Use:    "stop",
		Short:    "Stop a topology from an environment",
		Long:    "Stop a topology from an environment",
		RunE: invokeObjFunc("stop", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	stop.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	stop.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(stop)

	//restart
	var restart = &cobra.Command{
		Use:    "restart",
		Short:    "Restart a topology from an environment",
		Long:    "Restart a topology from an environment",
		RunE: invokeObjFunc("restart", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	restart.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	restart.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(restart)

	//setup
	var setup = &cobra.Command{
		Use:    "setup",
		Short:    "Setup configuration",
		Long:    "Setup configuration of all nodes in topology",
		RunE: invokeObjFunc("setup", []ArgType{{Flag:"id", Type:"uint"}}),
	}
	setup.Flags().UintVarP(&flag_id, "id", "i", 0, "Topology id")
	setup.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	topo.AddCommand(setup)
}

func topo(cmd *cobra.Command, args []string) error {
	return nil
}







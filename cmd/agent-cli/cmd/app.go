package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/cli/format"
	"github.com/ChaosXu/nerv/lib/cli"
)

func init() {
	var app = &cobra.Command{
		Use:    "app [command] [flags]",
		Aliases: []string{"App"},
		Short:    "Manage the app",
		Long:    "Manage the app",
		RunE: app,
	}
	RootCmd.AddCommand(app)

	//list
	var list = &cobra.Command{
		Use:    "list",
		Short:    "List all apps",
		Long:    "List all apps",
		RunE: cli.ListObjsFunc("App",
			&format.Page{List:"data", Columns:[]format.Column{
				{Name:"ID", Format:"%v"},
				{Name:"name", Label:"Name", Format:"%s"},
				{Name:"version", Label:"Version", Format:"%v"},
			}}),
	}
	list.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	app.AddCommand(list)

	//get
	var get = &cobra.Command{
		Use:    "get",
		Short:    "Get an app",
		Long:    "Get an app",
		RunE: cli.GetObjFunc("App", []string{}),
	}
	get.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "App id")
	get.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	app.AddCommand(get)


	//create
	var create = &cobra.Command{
		Use:    "create",
		Short:    "Create an app",
		Long:    "Create an app",
		RunE: cli.CreateObjFunc("App"),
	}
	create.Flags().StringVarP(&cli.Flag_data_path, "Data", "D", "", "required. The path of data file")
	create.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	app.AddCommand(create)



	//delete
	var delete = &cobra.Command{
		Use:    "delete",
		Short:    "Delete an app",
		Long:    "Delete an app",
		RunE: cli.RemoveObjFunc("App"),
	}
	delete.Flags().UintVarP(&cli.Flag_id, "id", "i", 0, "App id")
	delete.Flags().StringVarP(&cli.Flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	app.AddCommand(delete)




}

func app(cmd *cobra.Command, args []string) error {
	return nil
}







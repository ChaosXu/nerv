package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	libsvc "github.com/ChaosXu/nerv/lib/service"
	"github.com/ChaosXu/nerv/cmd/agent/service"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
)

var RootCmd = &cobra.Command{Use: "agent"}

func init() {

	//start
	var start = &cobra.Command{
		Use:    "start",
		Short:    "Start agent",
		Long:    "Start agent",
		RunE: serviceInit,
	}
	start.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	RootCmd.AddCommand(start)
}

func serviceInit(cmd *cobra.Command, args []string) error {
	env.InitByConfig(flag_config)

	container := libsvc.NewContainer()
	container.Add(&service.DBService{},"DB",nil)
	container.Add(&service.Agent{}, "Agent", &service.RemoteScriptServiceFactory{})
	container.Add(&service.AppService{}, "App", nil)
	container.Add(&service.HttpService{}, "HTTP", nil)
	container.Add(&rest.RestController{}, "RestController", nil)
	container.Build()
	defer container.Dispose()
	select {}
	return nil
}






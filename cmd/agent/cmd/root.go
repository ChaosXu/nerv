package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/env"
	_ "github.com/ChaosXu/nerv/cmd/agent/service"
	libsvc "github.com/ChaosXu/nerv/lib/service"
)

var RootCmd = &cobra.Command{Use: "agent"}

func init() {

	//start
	var start = &cobra.Command{
		Use:    "start",
		Short:    "Start agent",
		Long:    "Start agent",
		RunE: agentStart,
	}
	start.Flags().StringVarP(&flag_config, "config", "c", "../config/config.json", "The path of config.json. Default is ../config/config.json ")
	RootCmd.AddCommand(start)
}

func agentStart(cmd *cobra.Command, args []string) error {
	env.InitByConfig(flag_config)

	for _, svc := range libsvc.Registry.Services {
		if err := svc.Init(); err != nil {
			return err
		}
	}
	select {}
	return nil
}





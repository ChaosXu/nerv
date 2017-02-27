package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/cmd/agent/service"
	"github.com/ChaosXu/nerv/lib/env"
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

	agent, err := service.NewAgent(env.Config())
	if err != nil {
		return err
	}
	if err := agent.Start(); err != nil {
		return err
	}
	return nil
}





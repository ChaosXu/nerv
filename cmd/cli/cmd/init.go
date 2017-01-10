package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
)

var init_flag_template string
var init_flag_topology_name string

func init() {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Install nerv in one host",
		Long:
		`
		Use the init command to install all components of nerv in one host
		`,
		RunE: install,
	}
	initCmd.Flags().StringVarP(&init_flag_template, "template", "t", "", "required. The path of template that used to install nerv")
	initCmd.Flags().StringVarP(&init_flag_topology_name, "topolgoy", "o", "", "required. Topology name")

	RootCmd.AddCommand(initCmd)
}

func install(cmd *cobra.Command, args []string) error {
	if init_flag_template == "" {
		return errors.New("--template -t is null")
	}

	if init_flag_topology_name == "" {
		return errors.New("--topology -o is null")
	}

	template, err := topology.GetLocalTemplate(init_flag_template)
	if err != nil {
		return err
	}

	topology := template.NewTopology("nerv")
	return topology.Install()
}



package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ChaosXu/nerv/lib/env"
	libsvc "github.com/ChaosXu/nerv/lib/service"
	_ "github.com/ChaosXu/nerv/cmd/agent/service"
	"os"
	"fmt"
	"github.com/ChaosXu/nerv/lib/db"
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
	if _, err := os.Stat("../data"); err != nil {
		if err := os.MkdirAll("../data", os.ModeDir | os.ModePerm); err != nil {
			return fmt.Errorf("create dir ../data failed. %s", err.Error())
		}
	}

	err := initDB()
	if err != nil {
		return err
	}
	defer db.DB.Close()

	for _, factory := range libsvc.Registry.Services {
		if err := factory.Init(); err != nil {
			return err
		}

		svc := factory.Get()
		if svc != nil {
			initializer, ok := svc.(libsvc.Initializer)
			if ok {
				if err := initializer.Init(); err != nil {
					return err
				}
			}
		}
	}
	select {}
	return nil
}

func initDB() error {
	gdb, err := gorm.Open("sqlite3", "../data/agent.db")
	if err != nil {
		return err
	}
	db.DB = gdb
	db.DB.LogMode(false)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
	return nil
}





package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/controller/db"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize",
	Long: `initialize command. For example:

database init: init database
node init:     init node`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
var initdbCmd = &cobra.Command{
	Use:   "db",
	Short: "db init",
	Long:  "db init cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(db.InitDB())
		db.AddDBGroup(db.Group{
			Name:       "admin",
			Admin:      "test",
			User:       "test",
			MaxCPU:     100,
			MaxMem:     10240000,
			MaxStorage: 1000000000,
		})
		db.AddDBUser(db.User{Name: "test", Pass: "test"})
		return nil
	},
}
var initNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node init",
	Long:  "node init cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Not implemented")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initdbCmd)
	initCmd.AddCommand(initNodeCmd)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/imacon/db"
	"github.com/yoneyan/vm_mgr/imacon/etc"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		etc.ConfigGet()
		db.InitDB()
		fmt.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

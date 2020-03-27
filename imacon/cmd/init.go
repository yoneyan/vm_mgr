package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/imacon/db"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitDB()
		fmt.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

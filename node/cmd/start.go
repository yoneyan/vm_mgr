package cmd

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/data"
	"github.com/yoneyan/vm_mgr/node/vm"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start client server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		vm.StartupProcess()
		data.Server()
		fmt.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

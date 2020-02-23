package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/node/manage"
	"github.com/yoneyan/vm_mgr/node/vm"
	"log"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Long: `test command. 
For example:
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var testRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart qemu (name: test)",
	Long:  `test1 command`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vm.Restart("test")

		return nil
	},
}
var testAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add test",
	Long:  "add test",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
var testStorageAddCmd = &cobra.Command{
	Use:   "sadd",
	Short: "storage add test",
	Long:  "storage add test",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := manage.Storage{
			Path:   "/home/yoneyan",
			Name:   "test",
			Format: "qcow2",
			Size:   100,
		}

		err := manage.CreateStorage(&s)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.AddCommand(testAddCmd)
	testCmd.AddCommand(testRestartCmd)
	testCmd.AddCommand(testStorageAddCmd)

}

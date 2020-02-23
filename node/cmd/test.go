package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/node/vm"
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
		vm.Restart("name")

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

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.AddCommand(testAddCmd)
	testCmd.AddCommand(testRestartCmd)

}

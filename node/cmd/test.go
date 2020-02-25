package cmd

import (
	"fmt"
	"github.com/mattn/go-pipeline"
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
		vm.VMRestart("test")

		return nil
	},
}
var testAddCmd = &cobra.Command{
	Use:   "1",
	Short: "1",
	Long:  "1",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := pipeline.CombinedOutput(
			[]string{"ps", "axf"},
			[]string{"grep", "/" + "a" + ".sock"},
			[]string{"grep", "qemu"},
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("%s", out)

		fmt.Println("-------")
		out, err = pipeline.CombinedOutput(
			[]string{"ps", "axf"},
			[]string{"grep", "/" + "c" + ".sock"},
			[]string{"grep", "qemu"},
		)
		if err != nil {
			fmt.Println(false)
		}
		fmt.Printf("%s", out)

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

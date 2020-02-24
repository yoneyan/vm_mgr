/*
Copyright © 2020 yoneyan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	"log"
	"strconv"
)

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "create: vm create ,delete: vm delete",
	Long: `This is vm create and delete command.
create is vm create. Also, delete is vm delete.`,
}

var vmCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create vm",
	Long: `VM create tool
For example:
vm create -n test -c 1 -m 1024 -p /home/yoneyan/test.qcow2 -s 1024 -N br100 -v 200 -C /home/yoneyan/Downloads/ubuntu-18.04.4-live-server-amd64.iso -M false
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stringArray := []string{"name", "storage_path", "cdrom", "vnet"}
		int64Array := []string{"cpu", "mem", "storage", "vnc"}
		var resultStringArray [4]string
		var resultInt64Array [4]int64
		for i, b := range stringArray {
			result, err := cmd.PersistentFlags().GetString(b)
			if err != nil {
				log.Fatalf("could not greet: %v", err)
				return nil
			}
			resultStringArray[i] = result
		}
		for i, b := range int64Array {
			result, err := cmd.PersistentFlags().GetInt64(b)
			if err != nil {
				log.Fatalf("could not greet: %v", err)
				return nil
			}
			resultInt64Array[i] = result
		}

		change, err := cmd.PersistentFlags().GetBool("change")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.CreateVM(resultStringArray[0], resultInt64Array[0], resultInt64Array[1], resultInt64Array[2], resultStringArray[1], resultStringArray[2], resultStringArray[3], resultInt64Array[3], change)
		fmt.Println("Process End")
		return nil
	},
}
var vmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete vm",
	Long:  "VM Delete tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		data.DeleteVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "vm start",
	Long:  "VM start tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		data.StartVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var vmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "vm stop",
	Long:  "VM stop tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		data.StopVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var vmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "vm get",
	Long:  "VM get tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var vmGetIDCmd = &cobra.Command{
	Use:   "id",
	Short: "get id",
	Long:  "VM get tool from vmid",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		data.GetVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}
var vmGetNameCmd = &cobra.Command{
	Use:   "name",
	Short: "get name",
	Long:  "VM get tool from name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result := args[0]
		data.GetVMName(result)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	vmCreateCmd.PersistentFlags().StringP("name", "n", "none", "vm name")
	vmCreateCmd.PersistentFlags().Int64P("cpu", "c", 0, "virtual cpu")
	vmCreateCmd.PersistentFlags().Int64P("mem", "m", 0, "virtual memory")
	vmCreateCmd.PersistentFlags().StringP("storage_path", "p", "none", "storage path")
	vmCreateCmd.PersistentFlags().Int64P("storage", "s", 0, "storage capacity")
	vmCreateCmd.PersistentFlags().StringP("cdrom", "C", "", "cdrom path")
	vmCreateCmd.PersistentFlags().StringP("vnet", "N", "none", "virtual net")
	vmCreateCmd.PersistentFlags().Int64P("vnc", "v", 0, "vnc port")
	vmCreateCmd.PersistentFlags().BoolP("change", "M", false, "change spec")

	rootCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(vmCreateCmd)
	vmCmd.AddCommand(vmDeleteCmd)
	vmCmd.AddCommand(vmStartCmd)
	vmCmd.AddCommand(vmStopCmd)
	vmCmd.AddCommand(vmGetCmd)
	vmGetCmd.AddCommand(vmGetIDCmd)
	vmGetCmd.AddCommand(vmGetNameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

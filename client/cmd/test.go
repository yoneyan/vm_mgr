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
	"github.com/yoneyan/vm_mgr/client/data/direct"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "create: test create ,delete: test delete",
	Long: `This is test create and delete command.
create is test create. Also, delete is test delete.`,
}

var testCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create vm",
	Long: `VM create tool
For example:
vm create -n test -c 1 -m 1024 -p /home/yoneyan/test.qcow2 -s 1024 -N br100 -v 200 -C /home/yoneyan/Downloads/ubuntu-18.04.4-live-server-amd64.iso -a 1
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

		_, err := cmd.PersistentFlags().GetBool("autostart")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		c := pb.VMData{
			Vmname:  resultStringArray[0],
			Vcpu:    resultInt64Array[0],
			Vmem:    resultInt64Array[1],
			Storage: resultInt64Array[2],
			Vnet:    resultStringArray[3],
		}
		if direct.CreateVM(&c) {
			fmt.Println("Process End")
		} else {
			fmt.Println("Process Failed")
		}
		return nil
	},
}
var testDeleteCmd = &cobra.Command{
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
		direct.DeleteVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testStartCmd = &cobra.Command{
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
		direct.StartVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testStopCmd = &cobra.Command{
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
		direct.StopVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "vm shutdown",
	Long:  "VM shutdown tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		direct.ShutdownVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "vm reset",
	Long:  "VM reset tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		direct.ResetVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "vm resume",
	Long:  "VM resume tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		direct.ResumeVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "vm pause",
	Long:  "VM pause tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}
		direct.PauseVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}

var testGetCmd = &cobra.Command{
	Use:   "get",
	Short: "vm get",
	Long:  "VM get tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var testGetIDCmd = &cobra.Command{
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
		direct.GetVM(int64(result))
		fmt.Println("Process End")
		return nil
	},
}
var testGetNameCmd = &cobra.Command{
	Use:   "name",
	Short: "get name",
	Long:  "VM get tool from name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result := args[0]
		direct.GetVMName(result)
		fmt.Println("Process End")
		return nil
	},
}
var testGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "all",
	Long:  "get all VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		direct.GetAllVM(1)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	testCreateCmd.PersistentFlags().StringP("name", "n", "none", "vm name")
	testCreateCmd.PersistentFlags().Int64P("cpu", "c", 0, "virtual cpu")
	testCreateCmd.PersistentFlags().Int64P("mem", "m", 0, "virtual memory")
	testCreateCmd.PersistentFlags().StringP("storage_path", "p", "none", "storage path")
	testCreateCmd.PersistentFlags().Int64P("storage", "s", 0, "storage capacity")
	testCreateCmd.PersistentFlags().StringP("cdrom", "C", "", "cdrom path")
	testCreateCmd.PersistentFlags().StringP("vnet", "N", "none", "virtual net")
	testCreateCmd.PersistentFlags().Int64P("vnc", "v", 0, "vnc port")
	testCreateCmd.PersistentFlags().BoolP("autostart", "a", false, "autostart")

	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(testCreateCmd)
	testCmd.AddCommand(testDeleteCmd)
	testCmd.AddCommand(testStartCmd)
	testCmd.AddCommand(testStopCmd)
	testCmd.AddCommand(testShutdownCmd)
	testCmd.AddCommand(testResetCmd)
	testCmd.AddCommand(testPauseCmd)
	testCmd.AddCommand(testResumeCmd)
	testCmd.AddCommand(testGetCmd)

	testGetCmd.AddCommand(testGetIDCmd)
	testGetCmd.AddCommand(testGetNameCmd)
	testGetCmd.AddCommand(testGetAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

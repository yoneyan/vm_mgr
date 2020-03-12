package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	"github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
)

// vmDirectCmd represents the vm command
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
vm create -n test -c 1 -m 1024 -p /home/yoneyan/test.qcow2 -s 1024 -N br100 -v 200 -C /home/yoneyan/Downloads/ubuntu-18.04.4-live-server-amd64.iso -a 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stringArray := []string{"name", "storage_path", "cdrom", "vnet"}
		int64Array := []string{"cpu", "mem", "storage", "vnc", "node"}
		var resultStringArray [4]string
		var resultInt64Array [5]int64
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

		autostart, err := cmd.PersistentFlags().GetBool("autostart")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		d := Base(cmd)

		c := grpc.VMData{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Vmname:  resultStringArray[0],
			Node:    int32(resultInt64Array[4]),
			Vcpu:    resultInt64Array[0],
			Vmem:    resultInt64Array[1],
			Storage: resultInt64Array[2],
			Vnet:    resultStringArray[3],
			Option: &grpc.Option{
				StoragePath: resultStringArray[1],
				CdromPath:   resultStringArray[2],
				Vnc:         int32(resultInt64Array[3]),
				Autostart:   autostart,
			},
		}
		data.CreateVM(&c, d.Host)
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.DeleteVM(&c, d.Host)
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.StartVM(&c, d.Host)
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

		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.StopVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var vmShutdownCmd = &cobra.Command{
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

		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.ShutdownVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var vmResetCmd = &cobra.Command{
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.ResetVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var vmResumeCmd = &cobra.Command{
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.ResumeVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var vmPauseCmd = &cobra.Command{
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.PauseVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var vmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "vm get",
	Long:  "VM get tool",
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
		d := Base(cmd)
		c := grpc.VMID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Id: int64(result),
		}
		data.GetVM(&c, d.Host)
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
		d := Base(cmd)
		c := grpc.VMName{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Vmname: args[0],
		}
		data.GetVMName(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}
var vmGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "all",
	Long:  "get all VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		c := grpc.Base{
			User:  d.User,
			Pass:  d.Pass,
			Group: d.Group,
			Token: d.Token,
		}
		data.GetAllVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	vmCreateCmd.PersistentFlags().Int64P("node", "r", 1, "nodeid")
	vmCreateCmd.PersistentFlags().StringP("name", "n", "", "vm name")
	vmCreateCmd.PersistentFlags().Int64P("cpu", "c", 1, "virtual cpu")
	vmCreateCmd.PersistentFlags().Int64P("mem", "m", 512, "virtual memory")
	vmCreateCmd.PersistentFlags().StringP("storage_path", "P", "", "storage path")
	vmCreateCmd.PersistentFlags().Int64P("storage", "s", 1024, "storage capacity")
	vmCreateCmd.PersistentFlags().StringP("cdrom", "C", "", "cdrom path")
	vmCreateCmd.PersistentFlags().StringP("vnet", "N", "", "virtual net")
	vmCreateCmd.PersistentFlags().Int64P("vnc", "v", 0, "vnc port")
	vmCreateCmd.PersistentFlags().BoolP("autostart", "a", false, "autostart")

	rootCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(vmCreateCmd)
	vmCmd.AddCommand(vmDeleteCmd)
	vmCmd.AddCommand(vmStartCmd)
	vmCmd.AddCommand(vmStopCmd)
	vmCmd.AddCommand(vmShutdownCmd)
	vmCmd.AddCommand(vmResetCmd)
	vmCmd.AddCommand(vmPauseCmd)
	vmCmd.AddCommand(vmResumeCmd)
	vmCmd.AddCommand(vmGetCmd)

	vmGetCmd.AddCommand(vmGetIDCmd)
	vmGetCmd.AddCommand(vmGetNameCmd)
	vmGetCmd.AddCommand(vmGetAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmDirectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmDirectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

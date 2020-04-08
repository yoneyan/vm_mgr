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
//default connect (contorller user)
vm create -n te -c 1 -m 1024 -T 1 -s 10240 -i ubuntu,18.04 -a false -r 10 -H 127.0.0.1:50200 -u test -p test -g otaku
vm create -n te -c 1 -m 1024 -T 1 -s 1,10240,2,20480 -i ubuntu,18.04 -a false -r 10 -H 127.0.0.1:50200 -u test -p test -g otaku

//direct connect(node)
vm create -n te -c 1 -m 1024 -T 0 -P 1,/home/yoneyan,1,home/yoneyan -s 10240,10240 -N 1,br0,br100 -v 200 -C windows.iso,virtio.iso -a false -r 10 -H 127.0.0.1:50100
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		storage, err := cmd.Flags().GetString("storage")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		storagepath, err := cmd.Flags().GetString("storage_path")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		cdrom, err := cmd.Flags().GetString("cdrom")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		vnet, err := cmd.Flags().GetString("vnet")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		cpu, err := cmd.Flags().GetInt64("cpu")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		mem, err := cmd.Flags().GetInt64("mem")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		vmtype, err := cmd.Flags().GetInt32("type")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		vnc, err := cmd.Flags().GetInt64("vnc")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		node, err := cmd.Flags().GetInt32("node")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		imagename, err := cmd.Flags().GetString("imagename")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		imagetag, err := cmd.Flags().GetString("imagetag")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		autostart, err := cmd.PersistentFlags().GetBool("autostart")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		d := Base(cmd)

		c := grpc.VMData{
			Base:   &grpc.Base{User: d.User, Pass: d.Pass, Group: d.Group, Token: d.Token},
			Vmname: name, Node: node, Vcpu: cpu, Vmem: mem, Storage: storage, Vnet: vnet, Type: vmtype,
			Option: &grpc.Option{StoragePath: storagepath,
				CdromPath: cdrom, Vnc: int32(vnc), Autostart: autostart,
			},
			Image: &grpc.Image{Name: imagename, Tag: imagetag},
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
var vmGetGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "get group",
	Long: `VM get tool from name.
For example:
vm get group -H 127.0.0.1:50200 -t [token] -g [Group]
vm get group -H 127.0.0.1:50200 -u test -p test -g [Group]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		c := grpc.Base{
			User:  d.User,
			Pass:  d.Pass,
			Group: d.Group,
			Token: d.Token,
		}
		data.GetGroupVM(&c, d.Host)
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
	vmCreateCmd.PersistentFlags().Int32P("node", "r", 1, "nodeid")
	vmCreateCmd.PersistentFlags().StringP("name", "n", "", "vm name")
	vmCreateCmd.PersistentFlags().Int64P("cpu", "c", 1, "virtual cpu")
	vmCreateCmd.PersistentFlags().Int64P("mem", "m", 512, "virtual memory")
	vmCreateCmd.PersistentFlags().StringP("storage_path", "P", "", "storage path")
	vmCreateCmd.PersistentFlags().StringP("storage", "s", "1024", "storage capacity")
	vmCreateCmd.PersistentFlags().Int32P("type", "T", 0, "type")
	vmCreateCmd.PersistentFlags().StringP("cdrom", "C", "", "cdrom path")
	vmCreateCmd.PersistentFlags().StringP("vnet", "N", "", "virtual net")
	vmCreateCmd.PersistentFlags().Int64P("vnc", "v", 0, "vnc port")
	vmCreateCmd.PersistentFlags().BoolP("autostart", "a", false, "autostart")
	vmCreateCmd.PersistentFlags().StringP("imagename", "I", "", "image name")
	vmCreateCmd.PersistentFlags().StringP("imagetag", "i", "", "image tag")

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
	vmGetCmd.AddCommand(vmGetGroupCmd)
	vmGetCmd.AddCommand(vmGetAllCmd)
}

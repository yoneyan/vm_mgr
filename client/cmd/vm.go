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
storagetype
0: Custom
1: Normal Disk 
2: SSD 
3: NVMe 
4: iSCSI
storagemode
0: Default
1: VirtIOMode 
->For example: (-N 0,br100,br200)
networkmode
0: Default
1: VirtIOMode
->For example: (-P 1,/home/yoneyan,0,/home/yoneyan)

For example:
//default connect (contorller)
vm create -n te -c 1 -m 1024 -t 1 -s 10240 -N br0 -v 200 -C ubuntu.iso -a false -r 10 -H 127.0.0.1:50200 -u test -p test -g otaku
//direct connect(node)
vm create -n te -c 1 -m 1024 -t 0 -P 1,/home/yoneyan,1,home/yoneyan -s 10240,10240 -N br0 -v 200 -C windows.iso,virtio.iso -a false -r 10 -H 127.0.0.1:50100
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
		storagetype, err := cmd.Flags().GetInt32("storagetype")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		vnc, err := cmd.Flags().GetInt64("vnc")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		node, err := cmd.Flags().GetInt64("node")
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
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			Vmname: name, Node: int32(node), Vcpu: cpu, Vmem: mem, Storage: storage, Vnet: vnet, Storagetype: storagetype,
			Option: &grpc.Option{StoragePath: storagepath,
				CdromPath: cdrom, Vnc: int32(vnc), Autostart: autostart,
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
	vmCreateCmd.PersistentFlags().Int64P("node", "r", 1, "nodeid")
	vmCreateCmd.PersistentFlags().StringP("name", "n", "", "vm name")
	vmCreateCmd.PersistentFlags().Int64P("cpu", "c", 1, "virtual cpu")
	vmCreateCmd.PersistentFlags().Int64P("mem", "m", 512, "virtual memory")
	vmCreateCmd.PersistentFlags().StringP("storage_path", "P", "", "storage path")
	vmCreateCmd.PersistentFlags().StringP("storage", "s", "1024", "storage capacity")
	vmCreateCmd.PersistentFlags().Int32P("storagetype", "T", 0, "storage capacity")
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
	vmGetCmd.AddCommand(vmGetGroupCmd)
	vmGetCmd.AddCommand(vmGetAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmDirectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmDirectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

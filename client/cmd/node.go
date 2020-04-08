package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	grpc "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
)

type NodeData struct {
	ID        int32
	MaxCPU    int32
	MaxMem    int32
	Storage   string
	Net       string
	OnlyAdmin bool
}

// nodeCmd represents the client command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node",
	Long:  `node tool.`,
}

var nodeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "node add",
	Long: `node add tool
For example:
node add [HostName] [IPandPort] [MaxCPU]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}

		d1 := Base(cmd)
		d2 := NodeSpec(cmd)

		c := grpc.NodeData{
			Base: &grpc.Base{
				User:  d1.User,
				Pass:  d1.Pass,
				Group: d1.Group,
				Token: d1.Token,
			},
			NodeID:    d2.ID,
			Hostname:  args[0],
			IP:        args[1],
			Path:      d2.Storage,
			OnlyAdmin: d2.OnlyAdmin,
			Sepc: &grpc.SpecData{
				Maxcpu: d2.MaxCPU,
				Maxmem: d2.MaxMem,
				Net:    d2.Net,
			},
		}
		data.NodeAdd(&c, d1.Host)
		fmt.Println("Process End")
		return nil
	},
}

var nodeRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "node remove",
	Long:  "client remove tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}

		d := Base(cmd)
		c := grpc.NodeID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			NodeID: int32(result),
		}
		data.NodeRemove(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var nodeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "node get",
	Long:  "node get tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		data.GetNode(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var nodeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "client stop",
	Long:  "client stop tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result, _ := strconv.Atoi(args[0])
		if result < 0 {
			return errors.New("value failed")
		}

		d := Base(cmd)
		c := grpc.NodeID{
			Base: &grpc.Base{
				User:  d.User,
				Pass:  d.Pass,
				Group: d.Group,
				Token: d.Token,
			},
			NodeID: int32(result),
		}
		data.NodeStopVM(&c, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

func NodeSpec(cmd *cobra.Command) NodeData {
	id, err := cmd.PersistentFlags().GetInt32("id")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	cpu, err := cmd.PersistentFlags().GetInt32("maxcpu")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	mem, err := cmd.PersistentFlags().GetInt32("maxmem")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	storagepath, err := cmd.PersistentFlags().GetString("storage_path")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	onlyadmin, err := cmd.PersistentFlags().GetBool("onlyadmin")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return NodeData{
		ID:        id,
		MaxCPU:    cpu,
		MaxMem:    mem,
		Storage:   storagepath,
		OnlyAdmin: onlyadmin,
	}
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeStopCmd)
	nodeCmd.AddCommand(nodeAddCmd)
	nodeCmd.AddCommand(nodeRemoveCmd)
	nodeCmd.AddCommand(nodeGetCmd)

	nodeAddCmd.PersistentFlags().Int32P("id", "i", 0, "node id")
	nodeAddCmd.PersistentFlags().Int32P("maxcpu", "c", 0, "max cpu")
	nodeAddCmd.PersistentFlags().Int32P("maxmem", "m", 0, "max memory")
	nodeAddCmd.PersistentFlags().StringP("storage_path", "P", "", "storage path")
	nodeAddCmd.PersistentFlags().BoolP("onlyadmin", "a", true, "onlyadmin")
}

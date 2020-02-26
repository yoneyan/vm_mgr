package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/controller/db"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node command",
	Long: `node command. 
For example:
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var nodeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add node",
	Long:  `node add command`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostname, err := cmd.PersistentFlags().GetString("hostname")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}
		ip, err := cmd.PersistentFlags().GetString("ip")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}
		grpc_port, err := cmd.PersistentFlags().GetInt("grpc_port")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}
		ssh_port, err := cmd.PersistentFlags().GetInt("ssh_port")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}
		user, err := cmd.PersistentFlags().GetString("user")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}
		pass, err := cmd.PersistentFlags().GetString("pass")
		if err != nil {
			fmt.Println("could not greet: %v", err)
			return nil
		}

		result := db.AddDBController(db.Controller{HostName: hostname, IP: ip, GRPCPort: grpc_port, SSHPort: ssh_port, User: user, Pass: pass})
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}
var nodeRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove node",
	Long:  "User pass test tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}

		result := db.DeleteDBController(args[0])
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.AddCommand(nodeAddCmd)
	nodeCmd.AddCommand(nodeRemoveCmd)

	nodeAddCmd.PersistentFlags().StringP("hostname", "n", "", "username")
	nodeAddCmd.PersistentFlags().StringP("ip", "i", "", "ip address")
	nodeAddCmd.PersistentFlags().IntP("grpc_port", "g", 0, "grpc port")
	nodeAddCmd.PersistentFlags().IntP("ssh_port", "s", 0, "ssh port")
	nodeAddCmd.PersistentFlags().StringP("user", "u", "", "username")
	nodeAddCmd.PersistentFlags().StringP("pass", "p", "", "user password")

}

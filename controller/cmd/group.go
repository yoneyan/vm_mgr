package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/controller/db"
	"log"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "group command",
	Long: `group command.
For example:

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add group",
	Long: `add group command.
For Example:
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		admin, err := cmd.PersistentFlags().GetString("admin")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		user, err := cmd.PersistentFlags().GetString("user")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		maxcpu, err := cmd.PersistentFlags().GetInt("maxcpu")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		maxmem, err := cmd.PersistentFlags().GetInt("maxmem")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		maxstorage, err := cmd.PersistentFlags().GetInt("maxstorage")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		result := db.AddDBGroup(db.VmGroup{Name: name, Admin: admin, User: user, MaxCPU: maxcpu, MaxMem: maxmem, MaxStorage: maxstorage})
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}
var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete group",
	Long: `delete group command.
For example:

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}

		result := db.DeleteDBGroup(args[0])
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)

	groupAddCmd.PersistentFlags().StringP("name", "n", "", "username")
	groupAddCmd.PersistentFlags().StringP("user", "u", "", "ip address")
	groupAddCmd.PersistentFlags().StringP("admin", "a", "", "grpc port")
	groupAddCmd.PersistentFlags().IntP("maxcpu", "c", 0, "ssh port")
	groupAddCmd.PersistentFlags().IntP("maxmem", "m", 0, "username")
	groupAddCmd.PersistentFlags().IntP("maxstorage", "s", 0, "user password")

	groupCmd.AddCommand(groupAddCmd)
	groupCmd.AddCommand(groupDeleteCmd)
}

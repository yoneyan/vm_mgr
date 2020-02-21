package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/controller/db"
	"log"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user create and delete.",
	Long:  `user create and delete.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create user",
	Long: `User create tool
		example: user create -n test -p test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		pass, err := cmd.PersistentFlags().GetString("pass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		result := db.AddDBUser(db.VmUser{Name: name, Pass: pass})
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}
var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete user",
	Long:  "User Delete tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires id")
		}
		result := db.DeleteDBUser(args[0])
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}
var userTestCmd = &cobra.Command{
	Use:   "passtest",
	Short: "user pass test",
	Long:  "User pass test tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		pass, err := cmd.PersistentFlags().GetString("pass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		result := db.TestPassDBUser(name, pass)
		fmt.Println(result)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	userCreateCmd.PersistentFlags().StringP("name", "n", "", "username")
	userCreateCmd.PersistentFlags().StringP("pass", "p", "", "user password")

	userTestCmd.PersistentFlags().StringP("name", "n", "", "username")
	userTestCmd.PersistentFlags().StringP("pass", "p", "", "user password")

	rootCmd.AddCommand(userCmd)

	userCmd.AddCommand(userCreateCmd)
	userCmd.AddCommand(userDeleteCmd)
	userCmd.AddCommand(userTestCmd)

}

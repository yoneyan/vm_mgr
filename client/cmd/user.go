package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	"log"
)

// userCmd represents the user command
var usercmd = &cobra.Command{
	Use:   "user",
	Short: "user",
	Long: `user command. For example:
user add test
user remove test`,
}

var useraddcmd = &cobra.Command{
	Use:   "add",
	Short: "user add",
	Long: `user add tool
for example:

user add admin test -H 127.0.0.1:50200 -u admin -p `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("false")
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		fmt.Println(host)
		fmt.Println(host)
		authuser, err := cmd.Flags().GetString("authuser")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authpass, err := cmd.Flags().GetString("authpass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.AddUser(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1])

		fmt.Println("Process End")
		return nil
	},
}

var userremovecmd = &cobra.Command{
	Use:   "remove",
	Short: "user remove",
	Long:  "user remove tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("false")
		//}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authuser, err := cmd.Flags().GetString("authuser")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authpass, err := cmd.Flags().GetString("authpass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.RemoveUser(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0])

		fmt.Println("Process End")
		return nil
	},
}

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for user",
	Long:  "get tool for user",
}

var userGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "user get all",
	Long: `get all user
for example:
user get all -u test -p test -H 127.0.0.1:50200`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("false")
		//}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authuser, err := cmd.Flags().GetString("authuser")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authpass, err := cmd.Flags().GetString("authpass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.GetAllUser(&data.AuthData{Name: authuser, Pass1: authpass}, host)

		fmt.Println("Process End")
		return nil
	},
}

var userchangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for user",
	Long:  "change tool for user",
}

var userpasschangeCmd = &cobra.Command{
	Use:   "pass",
	Short: "change pass",
	Long: `change pass tool for user
for example:
user change pass [username] [newpass]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("false")
		//}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authuser, err := cmd.Flags().GetString("authuser")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authpass, err := cmd.Flags().GetString("authpass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.UserNameChange(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1])

		fmt.Println("Process End")
		return nil
	},
}

var usernamechangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for user
for example:
user change pass [before username] [after username]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("false")
		//}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authuser, err := cmd.Flags().GetString("authuser")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
		authpass, err := cmd.Flags().GetString("authpass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}

		data.UserPassChange(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1])

		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(usercmd)
	usercmd.AddCommand(useraddcmd)
	usercmd.AddCommand(userremovecmd)
	usercmd.AddCommand(userGetCmd)
	usercmd.AddCommand(userchangeCmd)

	userGetCmd.AddCommand(userGetAllCmd)
	userchangeCmd.AddCommand(usernamechangeCmd)
	userchangeCmd.AddCommand(userpasschangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

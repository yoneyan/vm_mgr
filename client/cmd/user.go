package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
)

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

		d := Base(cmd)

		data.AddUser(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1])

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
		d := Base(cmd)

		data.RemoveUser(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0])

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
		d := Base(cmd)

		data.GetAllUser(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host)

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
		if len(args) < 1 {
			return errors.New("false")
		}
		d := Base(cmd)

		data.UserNameChange(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1])

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
		if len(args) < 1 {
			return errors.New("false")
		}
		d := Base(cmd)

		data.UserPassChange(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1])

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
}

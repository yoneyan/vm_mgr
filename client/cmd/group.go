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
	"github.com/yoneyan/vm_mgr/client/data"
	"log"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "group command",
	Long: `group command
For example:

`,
}

var groupaddCmd = &cobra.Command{
	Use:   "add",
	Short: "group add",
	Long: `group add tool
for example:

group add otaku -H 127.0.0.1:50200 -u admin -p
group add [GroupName] [Network] [MaxVM] [MaxCPU] [MaxMem] [MaxStorage]`,
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

		data.AddGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1], args[2], args[3], args[4], args[5])

		fmt.Println("Process End")
		return nil
	},
}

var groupremoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "group remove",
	Long: `group remove tool
for example:

group remove otaku -H 127.0.0.1:50200 -u test -p test
group remove [GroupName]`,
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

		data.RemoveGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0])

		fmt.Println("Process End")
		return nil
	},
}

var groupgetCmd = &cobra.Command{
	Use:   "get",
	Short: "group remove",
	Long: `group remove tool
for example:`,
}

var groupgetallCmd = &cobra.Command{
	Use:   "all",
	Short: "group get all",
	Long: `group get tool
for example:

group get all -H 127.0.0.1:50200 -u test -p test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return nil
		}
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

		data.GetAllGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host)

		fmt.Println("Process End")
		return nil
	},
}

var groupgeteachCmd = &cobra.Command{
	Use:   "select",
	Short: "group get each group",
	Long: `group get tool
for example:
`,
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

		data.GetSelectGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0])

		fmt.Println("Process End")
		return nil
	},
}

var groupgetmyCmd = &cobra.Command{
	Use:   "my",
	Short: "group get my group",
	Long: `group get tool
for example:
`,
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

		data.GetMyGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host)

		fmt.Println("Process End")
		return nil
	},
}

var groupjoinCmd = &cobra.Command{
	Use:   "join",
	Short: "group join",
	Long: `group join tool
for example:`,
}

var groupjoinAddCmd = &cobra.Command{
	Use:   "add",
	Short: "group join add",
	Long: `group join add tool
for example:
group join add [Admin/User] [GroupName] [User]`,
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

		data.JoinAddGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1], args[2])

		fmt.Println("Process End")
		return nil
	},
}

var groupjoinRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "group join remove",
	Long: `group join remove tool
for example:
group join remove [Admin/User] [GroupName] [User]`,
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

		data.JoinRemoveGroup(&data.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1], args[2])

		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupaddCmd)
	groupCmd.AddCommand(groupremoveCmd)
	groupCmd.AddCommand(groupgetCmd)
	groupgetCmd.AddCommand(groupgetallCmd)
	groupgetCmd.AddCommand(groupgeteachCmd)
	groupgetCmd.AddCommand(groupgetmyCmd)
	groupCmd.AddCommand(groupjoinCmd)
	groupjoinCmd.AddCommand(groupjoinAddCmd)
	groupjoinCmd.AddCommand(groupjoinRemoveCmd)

	//groupCmd.PersistentFlags().StringP("host", "H", "127.0.0.1:50200", "host example: 127.0.0.1:50001")
	//groupCmd.PersistentFlags().StringP("authuser", "u", "test", "username")
	//groupCmd.PersistentFlags().StringP("authpass", "p", "test", "password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

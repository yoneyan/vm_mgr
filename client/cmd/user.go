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
	pb "github.com/yoneyan/vm_mgr/client/data/controller"
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

		pb.AddUser(&pb.AuthData{Name: authuser, Pass1: authpass}, host, args[0], args[1])

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
		fmt.Println("sa")
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

		pb.RemoveUser(&pb.AuthData{Name: authuser, Pass1: authpass}, host, args[0])

		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(usercmd)
	usercmd.AddCommand(useraddcmd)
	usercmd.AddCommand(userremovecmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

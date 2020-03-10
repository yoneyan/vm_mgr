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
	grpc "github.com/yoneyan/vm_mgr/proto/proto-go"
	"strconv"
)

// nodeCmd represents the client command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node",
	Long:  `node tool.`,
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
				User:  d[1],
				Pass:  d[2],
				Group: d[3],
			},
			NodeID: int32(result),
		}
		data.NodeStopVM(&c, d[0])
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeStopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

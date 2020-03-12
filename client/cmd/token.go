package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
)

// testCmd represents the test command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "create: token create ,delete: token delete",
	Long: `For example:
`,
}
var tokenGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "token generate",
	Long: `token generate tool
for example:

token generate -H 127.0.0.1:50200 -u test -p test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)

		data.GenerateToken(d.Host, d.User, d.Pass)

		fmt.Println("Process End")
		return nil
	},
}

var tokenRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "token remove",
	Long: `token remove tool
for example:

token remove [token] -H 127.0.0.1:50200`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || 2 < len(args) {
			return errors.New("false")
		}
		d := Base(cmd)

		data.DeleteToken(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0])

		fmt.Println("Process End")
		return nil
	},
}

var tokenGetAllCmd = &cobra.Command{
	Use:   "get",
	Short: "token all get",
	Long: `token get tool
for example:

token get -H 127.0.0.1:50200 -u admin -p test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {

			return errors.New("false")
		}
		d := Base(cmd)

		data.GetAllToken(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(tokenGenerateCmd)
	tokenCmd.AddCommand(tokenRemoveCmd)
	tokenCmd.AddCommand(tokenGetAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

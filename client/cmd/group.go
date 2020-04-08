package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
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

		d := Base(cmd)

		data.AddGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1], args[2], args[3], args[4], args[5])

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

		d := Base(cmd)

		data.RemoveGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0])

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
		d := Base(cmd)

		data.GetAllGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host)

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
		d := Base(cmd)

		data.GetSelectGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0])

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
		d := Base(cmd)

		data.GetMyGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host)

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
		d := Base(cmd)

		data.JoinAddGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1], args[2])

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
		d := Base(cmd)

		data.JoinRemoveGroup(&data.AuthData{Name: d.User, Pass: d.Pass, Token: d.Token}, d.Host, args[0], args[1], args[2])

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
}

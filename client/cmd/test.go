package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "create: test create ,delete: test delete",
	Long: `This is test create and delete command.
create is test create. Also, delete is test delete.`,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

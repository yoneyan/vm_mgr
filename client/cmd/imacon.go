package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	grpc "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
)

type ImaconData struct {
	ID       int32
	HostName string
	IP       string
	Status   int32
}

// imaconCmd represents the Imaocn command
var imaconcmd = &cobra.Command{
	Use:   "imacon",
	Short: "imacon",
	Long:  `imacon command. `,
}

var imaconaddcmd = &cobra.Command{
	Use:   "add",
	Short: "imacon add",
	Long: `imacon add tool
for example:

imacon add -I 1 -i 127.0.0.1:50300 -n imacon1 -H 127.0.0.1:50200 -u test -p test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Imaocn(cmd)

		data.AddImacon(&grpc.ImaconData{
			Base: &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group},
			Id:   i.ID, Hostname: i.HostName, IP: i.IP, Status: i.Status}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

var imaconremovecmd = &cobra.Command{
	Use:   "remove",
	Short: "imacon remove",
	Long: `imacon remove tool
for example:

imacon remove -I 1 -H 127.0.0.1:50200 -u test -p test `,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Imaocn(cmd)

		data.RemoveImacon(&grpc.ImaconData{Id: i.ID, Base: &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group}}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

var imaconGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for Imaocn",
	Long:  "get tool for Imaocn",
}

var imaconGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Imaocn get all",
	Long: `get all Imaocn
This command is only controller.
for example:

imaocn get all -u test -p test -H 127.0.0.1:50200`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)

		data.GetAllImacon(&grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group}, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imaconcmd)
	imaconcmd.AddCommand(imaconaddcmd)
	imaconcmd.AddCommand(imaconremovecmd)
	imaconcmd.AddCommand(imaconGetCmd)

	imaconGetCmd.AddCommand(imaconGetAllCmd)

	imaconcmd.PersistentFlags().Int32P("id", "I", 0, "id")
	imaconcmd.PersistentFlags().StringP("ip", "i", "", "Imaocn ip")
	imaconcmd.PersistentFlags().StringP("hostname", "n", "", "Imaocn hostname")
	imaconcmd.PersistentFlags().Int32P("status", "s", 0, "Imaocn status")

}

func Imaocn(cmd *cobra.Command) ImaconData {
	id, err := cmd.Flags().GetInt32("id")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	hostname, err := cmd.Flags().GetString("hostname")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	status, err := cmd.Flags().GetInt32("status")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return ImaconData{ID: id, HostName: hostname, IP: ip, Status: status}
}

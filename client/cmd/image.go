package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/client/data"
	grpc "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
)

type ImageData struct {
	ImaconID  int32
	ID        int32
	Name      string
	Tag       string
	Path      string
	Type      int32
	MinMem    int32
	Authority int32
	FileName  string
	Raddr     string
}

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "image",
	Long:  `image command. `,
}

var imageaddcmd = &cobra.Command{
	Use:   "add",
	Short: "image add",
	Long: `image add tool
---Type
0:AddISO 1:AddDISK 10:JoinISO 11:JoinDISK

for example:
image add -y [Type] -f [filename] -I [ImaconID] `,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Image(cmd)

		data.AddImage(&grpc.ImageData{
			Path: i.Path, Name: i.Name, Tag: i.Tag, Type: i.Type, Minmem: i.MinMem, Authority: i.Authority,
			Filename: i.FileName,
			Base:     &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group},
			Imaconid: i.ImaconID, Raddr: i.Raddr,
		}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

var imageremovecmd = &cobra.Command{
	Use:   "remove",
	Short: "image remove",
	Long:  "image remove tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Image(cmd)

		data.DeleteImage(&grpc.ImageData{
			Type: i.Type, Filename: i.FileName, Imaconid: i.ImaconID,
			Base: &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group},
		}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

var imageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for image",
	Long:  "get tool for image",
}

var imageGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "image get all",
	Long: `get all image
This command is only controller.
for example:

image get all -u test -p test -H 127.0.0.1:50200`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)

		data.GetAllImage(&grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group}, d.Host)
		fmt.Println("Process End")
		return nil
	},
}

var imageChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for image",
	Long:  "change tool for image",
}

var imageTagChangeCmd = &cobra.Command{
	Use:   "tag",
	Short: "change tag",
	Long: `change tag tool for image
for example:

image change tag -f [filename] -T [newtag] [auth...]
(controller)
image change tag -f [filename] -T [newtag] -I [imaconid] [auth...]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Image(cmd)

		data.ImageTagChange(&grpc.ImageData{
			Tag: i.Tag, Imaconid: i.ImaconID, Filename: i.FileName, Name: i.Name,
			Base: &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group},
		}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

var imageNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for image
for example:

(direct)
image change tag -f [filename] -N [newname] [auth...]
(controller)
image change tag -f [filename] -N [newname] -I [imaconid] [auth...]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := Base(cmd)
		i := Image(cmd)

		data.ImageNameChange(&grpc.ImageData{
			Imaconid: i.ImaconID, Filename: i.FileName, Name: i.Name,
			Base: &grpc.Base{User: d.User, Pass: d.Pass, Token: d.Group},
		}, d.Host)

		fmt.Println("Process End")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imageaddcmd)
	imageCmd.AddCommand(imageremovecmd)
	imageCmd.AddCommand(imageGetCmd)
	imageCmd.AddCommand(imageChangeCmd)

	imageGetCmd.AddCommand(imageGetAllCmd)
	imageChangeCmd.AddCommand(imageNameChangeCmd)
	imageChangeCmd.AddCommand(imageTagChangeCmd)

	imageCmd.PersistentFlags().Int32P("imaconid", "I", 0, "imaconid")
	imageCmd.PersistentFlags().Int32P("id", "i", 0, "id")
	imageCmd.PersistentFlags().StringP("filename", "f", "", "filename")
	imageCmd.PersistentFlags().StringP("name", "N", "", "name")
	imageCmd.PersistentFlags().StringP("tag", "T", "", "tag")
	imageCmd.PersistentFlags().Int32P("type", "y", 0, "image type")
	imageCmd.PersistentFlags().StringP("path", "P", "", "path")
	imageCmd.PersistentFlags().StringP("raddr", "r", "", "remoteaddress")
	imageCmd.PersistentFlags().Int32P("minmem", "m", 0, "minimum memory")
	imageCmd.PersistentFlags().Int32P("authority", "a", 0, "authority")

}

func Image(cmd *cobra.Command) ImageData {
	id, err := cmd.Flags().GetInt32("id")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	filename, err := cmd.Flags().GetString("filename")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	imagetype, err := cmd.Flags().GetInt32("type")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	tag, err := cmd.Flags().GetString("tag")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	raddr, err := cmd.Flags().GetString("raddr")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	minmem, err := cmd.Flags().GetInt32("minmem")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	authority, err := cmd.Flags().GetInt32("authority")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	imaconid, err := cmd.Flags().GetInt32("imaconid")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return ImageData{
		ImaconID: imaconid, ID: id, Name: name, Tag: tag, Path: path, Type: imagetype,
		MinMem: minmem, Authority: authority, FileName: filename, Raddr: raddr,
	}
}

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	command := []string{"--name ", "--vcpus ", "--ram ", "--os-type ", "--os-variant ", "--disk path= ", "--cdrom ", "--network=bridge:", "--graphics vnc,listen=0.0.0.0,port="}
	value := []string{"none", "none", "none", "none", "none", "none", "none", "none", "none", "none"}

	var path string
	var size int

	app := cli.NewApp()

	app.Name = "vm_mgr"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.3"
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"c"},
			Usage:   "install kvm",
			Action: func(c *cli.Context) error {
				if f, err := os.Stat("/etc/redhat-release"); os.IsNotExist(err) || f.IsDir() {
					fmt.Println("Not exists！")
				} else {
					fmt.Println("This working os is CentOS")
					out, err := exec.Command("sudo ", "yum", "-y", "install", "qemu-img", "qemu-kvm", "libvirt", "virt-install", "bridge-utils").Output()
					if err != nil {
						fmt.Println(err.Error())
						os.Exit(1)
					}
					fmt.Println(string(out))
				}

				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "vm create command",
			Subcommands: []cli.Command{
				{
					Name:  "storage",
					Usage: "storage",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:        "size, s",
							Usage:       "disk size",
							Destination: &size,
						},
						cli.StringFlag{
							Name:        "path, p",
							Usage:       "disk path",
							Destination: &path,
						},
					},
					Action: func(c *cli.Context) error {
						out, err := exec.Command("qemu-img", "create", "-f", "qcow2", path, strconv.Itoa(size)).Output()
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						fmt.Println(string(out))

						return nil
					},
				},
				{
					Name:  "vm",
					Usage: "vm",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "vm name",
							Destination: &value[0],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "c",
							Required:    true,
							Destination: &value[1],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "m",
							Required:    true,
							Destination: &value[2],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "t",
							Usage:       "os-type",
							Destination: &value[3],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "v",
							Usage:       "os-variant",
							Destination: &value[4],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "disk",
							Usage:       "disk path",
							Destination: &value[5],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "iso",
							Usage:       "iso",
							Destination: &value[6],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "net, N",
							Usage:       "network",
							Destination: &value[7],
							Value:       "none",
						},

						cli.StringFlag{
							Name:        "vnc",
							Usage:       "vnc port",
							Destination: &value[8],
							Value:       "none",
						},
					},
					Action: func(c *cli.Context) error {

						for i := range command {
							fmt.Printf("%s: %s\n", command[i], value[i])
						}

						var command_exec []string

						for i := range value {
							if value[i] == "none" {

							} else {
								command_exec = append(command_exec, command[i]+value[i])
							}
						}

						fmt.Println(command_exec)

						out, err := exec.Command("virt-install", command_exec...).Output()
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
						fmt.Println(string(out))

						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

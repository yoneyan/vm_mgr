package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

func main() {

	command := []string{"--name ", "--vcpus ", "--memory ", "--os-type ", "--os-variant ", "--disk path=", "--cdrom ", "--network=bridge:", "--graphics vnc,listen=0.0.0.0,port="}
	value := []string{"none", "none", "none", "none", "none", "none", "none", "none", "none", "none"}

	var path, size string

	app := cli.NewApp()

	app.Name = "vm_mgr"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.3.6.5"
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install kvm",
			Action: func(c *cli.Context) error {
				if f, err := os.Stat("/etc/redhat-release"); os.IsNotExist(err) || f.IsDir() {
					fmt.Println("not compatible")
				} else {
					fmt.Println("This working os is CentOS")
					out, err := exec.Command("yum", "-y", "install", "qemu-img", "qemu-kvm", "libvirt", "virt-install", "bridge-utils").Output()
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
						cli.StringFlag{
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
						out, err := exec.Command("qemu-img", "create", "-f", "qcow2", path, size).Output()
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
							Name:        "n,name",
							Usage:       "vm name",
							Destination: &value[0],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "c,core",
							Required:    true,
							Destination: &value[1],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "m,memory",
							Required:    true,
							Destination: &value[2],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "t,type",
							Usage:       "os-type",
							Destination: &value[3],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "v,variant",
							Usage:       "os-variant",
							Destination: &value[4],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "d,disk",
							Usage:       "disk path",
							Destination: &value[5],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "I,iso",
							Usage:       "iso",
							Destination: &value[6],
							Value:       "none",
						},
						cli.StringFlag{
							Name:        "N,net",
							Usage:       "network",
							Destination: &value[7],
							Value:       "none",
						},

						cli.StringFlag{
							Name:        "V,vnc",
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
						/*
							command_exec = append(command_exec, "-c")
							command_exec = append(command_exec, "virt-install")
						*/
						for i := range value {
							if value[i] == "none" {

							} else {
								command_exec = append(command_exec, command[i]+value[i])
								/*command_exec = append(command_exec,command[i])
								command_exec = append(command_exec,value[i])*/
							}
						}

						fmt.Println(command_exec)

						cmd := exec.Command(os.Getenv("SHELL"), "-c", strings.Join(command_exec, " "))
						output, err := cmd.CombinedOutput()
						if err != nil {
							fmt.Println(fmt.Sprint(err) + ": " + string(output))
							return nil
						} else {
							fmt.Println(string(output))
						}

						/*
							out, err := exec.Command("echo", command_exec...).Output()
							fmt.Println(string(out))
							if err != nil {
								fmt.Println(err.Error() + ": " + string(out))
								os.Exit(1)
							}*/
						return nil

					},
				},
			},
		},
	}

	app.Run(os.Args)
}

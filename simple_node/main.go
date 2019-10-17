package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Name = "vm_mgr"
	app.Usage = "This app echo input arguments"
	app.Version = "0.1"
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
						cli.StringFlag{Name: "size, s"},
						cli.StringFlag{Name: "path, p"},
					},
					Action: func(c *cli.Context) error {
						out, err := exec.Command("qemu-img", "create", "-f", "qcow2", c.String("path"), c.String("size")).Output()
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
						cli.StringFlag{Name: "n,name"},
						cli.StringFlag{Name: "c,core"},
						cli.StringFlag{Name: "m,memory"},
						cli.StringFlag{Name: "t,type"},
						cli.StringFlag{Name: "v,variant"},
						cli.StringFlag{Name: "d,disk"},
						cli.StringFlag{Name: "I,iso"},
						cli.StringFlag{Name: "N,net"},
						cli.StringFlag{Name: "V,vnc"},
					},
					Action: func(c *cli.Context) error {
						var command_exec []string
						command_exec = append(command_exec, "virt-install")
						command_exec = append(command_exec, "--name "+c.String("name"))
						if c.String("core") != "" {
							command_exec = append(command_exec, "--vcpus"+c.String("core"))
						}
						if c.String("memory") != "" {
							command_exec = append(command_exec, "--memory"+c.String("memory"))
						}
						if c.String("os-type") != "" {
							command_exec = append(command_exec, "--os-type"+c.String("type"))
						}
						if c.String("variant") != "" {
							command_exec = append(command_exec, "--os-variant"+c.String("variant"))
						}
						if c.String("disk") != "" {
							command_exec = append(command_exec, "--disk path="+c.String("disk"))
						}
						if c.String("iso") != "" {
							command_exec = append(command_exec, "--cdrom "+c.String("iso"))
						}
						if c.String("net") != "" {
							command_exec = append(command_exec, "--network=bridge:"+c.String("net"))
						}
						if c.String("vnc") != "" {
							command_exec = append(command_exec, "--graphics vnc,listen=0.0.0.0,port="+c.String("vnc"))
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
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

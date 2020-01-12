package main

import (
	"fmt"
	"os"

	//"github.com/urfave/cli"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "vm_mgr"
	app.Usage = "v0.001"
	app.Version = "0.001"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "node add",
			Action: func(c *cli.Context) error {
				fmt.Println("Not implemented")
				return nil
			},
		},
		{
			Name:    "test-1",
			Aliases: []string{"t1"},
			Usage:   "test1 script",
			Action: func(c *cli.Context) error {
				fmt.Println(initdb())
				return nil
			},
		},
		{
			Name:    "test-2",
			Aliases: []string{"t2"},
			Usage:   "test2 script",
			Action: func(c *cli.Context) error {
				fmt.Println(sshTest())
				return nil
			},
		},
	}
	app.Run(os.Args)
}
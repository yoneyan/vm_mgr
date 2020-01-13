package main

import (
	"fmt"
	"os"
	"strconv"

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
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init",
			Subcommands: []cli.Command{
				{
					Name:  "db",
					Usage: "db init",
					Action: func(c *cli.Context) error {
						fmt.Println(initdb())
						return nil
					},
				},
				{
					Name: "node",
					Usage: "node",
					Action: func(c *cli.Context) error{
						fmt.Println("Not implemented")
						return nil
					},
				},
			},
		},
		{
			Name:    "vm",
			Aliases: []string{"v"},
			Usage:   "vm",
			Action: func(c *cli.Context) error {
				fmt.Println("Not implemented")
				return nil
			},
		},
		{
			Name:    "node",
			Aliases: []string{"n"},
			Usage:   "node",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "node add",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "hostname, H"},
						cli.StringFlag{Name: "ip, i"},
						cli.StringFlag{Name: "port, p"},
						cli.StringFlag{Name: "user, u"},
						cli.StringFlag{Name: "pass, P"},
					},
					Action: func(c *cli.Context) error {
						var port int
						port, _ = strconv.Atoi(c.String("port"))
						result := sshTest(c.String("ip"), port, c.String("user"), c.String("pass"))
						fmt.Println(result)
						result = db_controller("add", c.String("hostname"), c.String("ip"), port, c.String("user"), c.String("pass"))
						fmt.Println(result)

						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "node remove",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "hostname, H"},
					},
					Action: func(c *cli.Context) error {
						zero := "zero"
						result := db_controller("remove", c.String("hostname"), zero, 0, zero, zero)
						fmt.Println(result)

						return nil
					},
				},
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
				fmt.Println("test")
				return nil
			},
		},
	}
	app.Run(os.Args)
}

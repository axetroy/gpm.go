package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"os/exec"
	"github.com/axetroy/gpm.go/lib"
)

func main() {
	app := cli.NewApp()
	app.Name = "gpm"
	app.Usage = "Git Package Manager, make you manage the repository easier, Power by Go"

	var config gpm.ConfigS= gpm.GetConfig()
	fmt.Println(config)
	fmt.Println(config.Paths.Root)

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a repository to gpm",
			Action:  func(c *cli.Context) error {
				fmt.Println("Repo: ", c.Args().First())
				var repo string = c.Args().First()
				cmd := exec.Command("git", "clone", repo)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if e := cmd.Start(); nil != e {
					fmt.Printf("ERROR: %v\n", e)
				}
				if e := cmd.Wait(); nil != e {
					fmt.Printf("ERROR: %v\n", e)
				}
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Remove a repository from registry and disk",
			Action:  func(c *cli.Context) error {
				fmt.Println("remove repo: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"c"},
			Usage:   "Display the all repositories in registry",
			Action:  func(c *cli.Context) error {
				fmt.Println("get list: ", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
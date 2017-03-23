package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
	"github.com/gpmer/gpm.go/lib"
	"os"
	"os/exec"
	"path"
)

func main() {

	gpm.Prepare()
	var config gpm.ConfigS = gpm.GetConfig()

	commander.Program.
		Command("gpm").
		Description("Git Package Manager, make you manage the repository easier, Power by Go")

	commander.Program.
		Command("add <repo>").
		Aliases([]string{"a"}).
		Description("add a repository to gpm").
		Action(func(c commander.Context) {
			var repo string = c.GetString("<repo>")

			os.Chdir(config.Paths.Temp)

			cmd := exec.Command("git", "clone", repo)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if e := cmd.Start(); nil != e {
				fmt.Printf("ERROR: %v\n", e)
			}
			if e := cmd.Wait(); nil != e {
				fmt.Printf("ERROR: %v\n", e)
			}

			fmt.Println("clone repo to", path.Join(config.Paths.Temp))
		})

	commander.Program.
		Command("remove <repo>").
		Aliases([]string{"rm"}).
		Description("remove a repository from registry and disk").
		Action(func(c commander.Context) {
			var repo string = c.GetString("<repo>")
			fmt.Println("remove repo:", repo)
		})

	commander.Program.
		Command("list <repo>").
		Aliases([]string{"c"}).
		Description("display the all repositories in registry").
		Action(func(c commander.Context) {
			var repo string = c.GetString("<repo>")
			fmt.Println("get list:", repo)
		})

	commander.Program.Parse()
}

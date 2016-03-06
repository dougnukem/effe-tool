package main

import (
	"effe-tool/builder"
	"effe-tool/docker"
	"effe-tool/factory"
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	app := cli.NewApp()
	app.Name = "effe-tool"
	app.Usage = "Utility to create, build and use effes."
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create a new empty effe.",
			Action:  factory.CreateNewEffe,
		},
		{
			Name:    "compile",
			Aliases: []string{"c"},
			Usage:   "Compile a single file or a whole directory passed as argument.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dirout",
					Value: "out/",
					Usage: "Directory where to save the executables.",
				},
				cli.StringFlag{
					Name:  "out",
					Value: "",
					Usage: "Custom name to save your executable.",
				},
			},
			Action: builder.Compile,
		},
		{
			Name:    "docker",
			Aliases: []string{"d"},
			Usage:   "Create docker images of a single executable or of every executable in the directory passed as argument.",
			Action:  docker.Dockerify,
		},
	}

	app.Run(os.Args)
}

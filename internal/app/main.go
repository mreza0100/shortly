package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

type actions struct {
	cfg *Configs
}

func main() {
	app := cli.NewApp()
	actions := new(actions)
	actions.cfg = loadConfigs()

	app.Commands = []cli.Command{
		{
			Name:    "run server",
			Aliases: []string{"run"},
			Usage:   "run the server with the given config file",
			Action:  actions.run,
		}, {
			Name:    "seed database",
			Aliases: []string{"seed"},
			Usage:   "seed the database",
			Action:  actions.seed,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

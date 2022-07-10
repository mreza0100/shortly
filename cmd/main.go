package main

import (
	"log"
	"os"

	"github.com/mreza0100/shortly/cmd/providers"
	"github.com/urfave/cli"
)

type actions struct {
	cfg *providers.Configs
}

func main() {
	app := cli.NewApp()
	actions := new(actions)
	actions.cfg = providers.ConfigProvider()

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
		}, {
			Name:    "health check",
			Aliases: []string{"healthcheck"},
			Usage:   "check the health of the server",
			Action:  actions.healthCheck,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

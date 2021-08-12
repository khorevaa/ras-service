package main

import (
	"github.com/khorevaa/ras-service/cmd"

	"os"
	"time"

	"github.com/khorevaa/logos"
	"github.com/urfave/cli/v2"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var log = logos.New("github.com/khorevaa/ras-service").Sugar()

func main() {

	app := &cli.App{
		Name:     "ras-service",
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Aleksey Khorev",
			},
		},
		Usage:     "RAS service for server 1C.Enterprise",
		UsageText: "ras-service command [command options] [arguments...]",
		Copyright: "(c) 2021 Khorevaa",
		//Description: "Command line utilities for server 1S.Enterprise",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "master-node-url", Aliases: []string{"M"},
				Value: "", Usage: "master node url", EnvVars: []string{"RS_MASTER_NODE_URL"}},
			&cli.StringFlag{
				Name: "node-name", Aliases: []string{"N"},
				Value: "", Usage: "node name", EnvVars: []string{"RS_NODE_NAME"}, DefaultText: "Host name"},
		},
	}

	for _, command := range cmd.Commands {
		app.Commands = append(app.Commands, command.Cmd())
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

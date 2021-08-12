package cmd

import (
	"github.com/khorevaa/logos"
	"github.com/urfave/cli/v2"
)

var log = logos.New("github.com/khorevaa/ras-service/cmd")

var Commands = []Command{

	&svcCmd{},
	&remoteCmd{},
}

type Command interface {
	Cmd() *cli.Command
}

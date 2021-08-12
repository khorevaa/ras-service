package cmd

import (
	"github.com/khorevaa/logos"
	"github.com/urfave/cli/v2"
)

var log = logos.New("github.com/v8platform/onec-util")

var Commands = []Command{

	&joinCmd{},
	&disconnectCmd{},
	//&clusterCmd{
	//	sub: []Command{
	//		//&disconnectCmd{},
	//		&joinCmd{},
	//	},
	//},
}

type Command interface {
	Cmd() *cli.Command
}

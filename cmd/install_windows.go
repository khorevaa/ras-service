//go:build windows
// +build windows

package cmd

import (
	"github.com/urfave/cli/v2"
)

type installCmd struct {
	serviceName  string
	serviceID    string
	serviceDescr string
}

func (c *installCmd) Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:        "install",
		Usage:       "Install service control for RAS 1S.Enterprise",
		Description: "Выполняет регистрацию службы для контроля RAS",
		ArgsUsage:   "ARGS",
		Action:      c.run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &c.serviceID, Name: "id",
				Value: "ras-service", Usage: "service id service for register"},
			&cli.StringFlag{
				Destination: &c.serviceName, Name: "name",
				Value: "RAS Control Service", Usage: "name for service"},
			&cli.StringFlag{
				Destination: &c.serviceDescr, Name: "descr",
				Value: "RAS Control Service", Usage: "description for service", DefaultText: ""},
		},
	}

	return cmd
}

func (c *installCmd) run(ctx *cli.Context) error {

	args := ctx.Args().Slice()
	err := installService(c.serviceID, c.serviceName, c.serviceDescr, args)

	if err != nil {
		return err
	}

	return startService(c.serviceID, args)

}

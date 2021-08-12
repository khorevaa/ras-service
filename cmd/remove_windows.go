//go:build windows
// +build windows

package cmd

import "github.com/urfave/cli/v2"

type removeCmd struct {
	serviceID string
}

func (c *removeCmd) Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:        "remove",
		Usage:       "Remove service control for RAS 1S.Enterprise",
		Description: "Выполняет удаление службы для контроля RAS",
		Action:      c.run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &c.serviceID, Name: "id",
				Value: "ras-service", Usage: "service id service for remove"},
		},
	}

	return cmd
}

func (c *removeCmd) run(_ *cli.Context) error {

	err := stopService(c.serviceID)

	if err != nil {
		log.Warn(err.Error())
	}

	return removeService(c.serviceID)
}

package cmd

import (
	"fmt"
	"github.com/khorevaa/ras-service/service"
	"github.com/urfave/cli/v2"
	"net"
	"os"
	"time"
)

type svcCmd struct {
	workDir string

	host string
	port string

	rasPath   string
	v8version string

	masterNodeUrl string
	nodeName      string

	healthcheck time.Duration //

	usr, pwd string // Авторизация на агенте

}

func (c *svcCmd) run(ctx *cli.Context) error {

	// masterNodeUrl := ctx.String("master-node-url")
	// nodeName := ctx.String("node-name")

	srv := service.Control{
		Starter: service.NewRasStarter(c.rasPath, []string{"cluster", fmt.Sprintf("%s:%s", c.host, c.port)}),
	}

	return srv.Run(ctx.Context)

}

func (c *svcCmd) Cmd() *cli.Command {

	cmd := &cli.Command{
		Name:        "svc",
		Usage:       "Control control for RAS 1S.Enterprise",
		Description: "Выполняет запуск службы для контроля RAS",
		ArgsUsage:   "HOST[:PORT]",
		Action:      c.run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &c.workDir, Name: "work-dir", Aliases: []string{"D"},
				Value: ".", Usage: "wording directory", EnvVars: []string{"RS_WORK_DIR"}},
			&cli.StringFlag{
				Destination: &c.rasPath, Name: "ras-path",
				Value: "", Usage: "full path for ras.exe", EnvVars: []string{"RS_RAS_PATH"}},
			&cli.StringFlag{
				Destination: &c.v8version, Name: "v8-version",
				Value: "8.3", Usage: "v8 version for find ras path", EnvVars: []string{"RS_V8_VERSION"}},
			&cli.DurationFlag{
				Destination: &c.healthcheck, Name: "healthcheck-period",
				Value: 30 * time.Minute, Usage: "healthcheck period ras state", EnvVars: []string{"RS_HEALTHCKECK_PERIOD"}, DefaultText: "30m"},
			&cli.StringFlag{
				Destination: &c.usr, Name: "user", Aliases: []string{"U"},
				Value: "", Usage: "server agent user", EnvVars: []string{"RS_AGENT_USER"}, DefaultText: ""},
			&cli.StringFlag{
				Destination: &c.pwd, Name: "pwd", Aliases: []string{"P"},
				Value: "", Usage: "server agent password", EnvVars: []string{"RS_AGENT_PASSWORD"}, DefaultText: ""},
			&cli.StringFlag{
				Destination: &c.port, Name: "port", Aliases: []string{"p"},
				Value: "", Usage: "server port number", EnvVars: []string{"RS_HOST_PORT"}, DefaultText: ""},
		},

		Before: func(ctx *cli.Context) error {

			hostPort, _ := os.LookupEnv("RS_CLUSTER_HOST")

			if ctx.Args().Len() > 0 {
				hostPort = ctx.Args().First()
			}

			host, port, _ := net.SplitHostPort(hostPort)

			c.host = host

			if len(port) > 0 {
				c.port = port
			}

			return nil

		},
	}

	return cmd
}

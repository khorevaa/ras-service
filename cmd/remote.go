package cmd

import (
	"github.com/urfave/cli/v2"
	"net"
	"os"
	"time"
)

type remoteCmd struct {
	host string
	port string

	masterNodeUrl string
	nodeName      string

	healthcheck time.Duration //

	usr, pwd string // Авторизация на агенте

}

func (c *remoteCmd) run(context *cli.Context) error {

	log.Info("run")

	//
	// mng, err := rac.NewManager(c.host, rac.ManagerOptions{
	// 	Timeout:         2 * time.Second,
	// 	TryTimeoutCount: 3,
	// 	DetectCluster:   false,
	// })
	//
	// if err != nil {
	// 	return err
	// }
	//
	// c.cluster, err = getCluster(mng)
	//
	// if err != nil {
	// 	log.Error("get clusters", logos.Error(err))
	// 	return err
	// }
	//
	// log.Debug("get cluster id", logos.String("cluster-id", c.cluster))
	//
	// err = mng.SetDefCluster(c, c)
	// if err != nil {
	// 	log.Error("setup work cluster", logos.Error(err), logos.String("cluster-id", c.ClusterSig()))
	// 	return err
	// }
	//
	//
	// if err != nil {
	// 	return err
	// }
	// //pp.Println(servers)

	return nil
}

func (c *remoteCmd) Cmd() *cli.Command {

	cmd := &cli.Command{
		Name:        "service",
		Usage:       "Service control for RAS 1S.Enterprise",
		Description: "Выполняет запуск службы для контроля RAS",
		ArgsUsage:   "HOST[:PORT]",
		Action:      c.run,
		Flags: []cli.Flag{
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

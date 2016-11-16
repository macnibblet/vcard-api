package cmd

import (
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"fmt"
)

var (
	CMD_HTTP = &cli.Command{
		Name: "http",
		Usage: "Start the http server",
		Action: runHttpServer,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name: "port",
				Usage: "define the port used to the run the webserver",
				Value: 7000,
				EnvVars: []string{"PORT"},
			},

			cli.StringFlag{
				Name: "bind-address",
				Usage: "Used to bind the listening port to a specific ip address",
				Value: "",
			},

			FLAG_DB_DSN,
		},
	}
)

func runHttpServer(ctx *cli.Context) error {

	return http.ListenAndServe(fmt.Sprintf("%s:%s", ctx.String("bind-address"), ctx.Int("port")), nil)
}

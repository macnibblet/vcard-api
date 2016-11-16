package cmd

import (
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"fmt"
	"gopkg.in/pg.v5"
	"github.com/macnibblet/vcard-api/endpoints"
	"github.com/gorilla/mux"
)

var (
	CMD_HTTP = &cli.Command{
		Name: "http",
		Usage: "Start the http server",
		Action: runHttpServer,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "port",
				Usage: "define the port used to the run the webserver",
				Value: 7000,
				EnvVars: []string{"PORT"},
			},

			&cli.StringFlag{
				Name: "bind-address",
				Usage: "Used to bind the listening port to a specific ip address",
				Value: "",
			},

			&cli.StringFlag{
				Name: "static-folder",
				Usage: "Path to the vcards that is hosted statically",
				Value: "./cards",
			},

			FLAG_DB_DSN,
		},
	}
)

func runHttpServer(ctx *cli.Context) error {

	db := pg.Connect(&pg.Options{

	})

	router := mux.NewRouter()
	router.PathPrefix("/cards/").Handler(http.StripPrefix("/cards/", http.FileServer(http.Dir(ctx.String("static-folder")))))


	endpoints.InjectUserRoutes(router)

	return http.ListenAndServe(fmt.Sprintf("%s:%s", ctx.String("bind-address"), ctx.Int("port")), nil)
}

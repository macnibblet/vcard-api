package cmd

import (
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/macnibblet/vcard-api/cards"
	"github.com/macnibblet/vcard-api/database"
	"github.com/macnibblet/vcard-api/endpoints"
	"gopkg.in/pg.v5"
	"gopkg.in/urfave/cli.v2"
	"net/http"
)

var (
	CMD_HTTP = &cli.Command{
		Name:   "http",
		Usage:  "Start the http server",
		Action: runHttpServer,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "define the port used to the run the webserver",
				Value:   7000,
				EnvVars: []string{"PORT"},
			},

			&cli.StringFlag{
				Name:  "bind-address",
				Usage: "Used to bind the listening port to a specific ip address",
				Value: "",
			},

			&cli.StringFlag{
				Name:  "static-folder",
				Usage: "Path to the vcards that is hosted statically",
				Value: "./static",
			},

			&cli.StringFlag{
				Name:    "jwt-token",
				Usage:   "The jwt secret token",
				Value:   "helloWorld",
				EnvVars: []string{"JWT_TOKEN"},
			},

			FLAG_DB_DSN,
		},
	}
)

func runHttpServer(ctx *cli.Context) error {

	// todo: replace configuration with DB_DSN
	db := pg.Connect(&pg.Options{
		Network:  "tcp",
		Addr:     "localhost",
		User:     "vcard",
		Password: "vcard",
		Database: "vcard",
	})

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(ctx.String("jwt-token")), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

	context := &cards.Context{
		JwtMiddleware:  jwtMiddleware,
		UserRepository: &database.UserRepository{DB: db},
		CardRepository: &database.VCardRepository{DB: db},
	}

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(ctx.String("static-folder")))))

	endpoints.InjectUserRoutes(context, router)

	fmt.Println("Starting webserver...")
	
	return http.ListenAndServe(fmt.Sprintf("%s:%d", ctx.String("bind-address"), ctx.Int("port")), nil)
}

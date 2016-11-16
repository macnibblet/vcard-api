package cmd

import (
	"github.com/AndrewVos/mig"
	"gopkg.in/urfave/cli.v2"
)

var (
	FLAG_DB_DSN = cli.StringFlag{
		Name:    "database-dsn",
		Value:   "host=localhost user=vcard password=vcard dbname=vcard sslmode=disable",
		EnvVars: []string{"DATABASE_URL"}, // This is for heroku
	}

	CMD_MIGRATE = &cli.Command{
		Name:   "migrate",
		Usage:  "Migrate the database to the latest version",
		Action: migrate,
		Flags:  []cli.Flag{FLAG_DB_DSN},
	}
)

func migrate(c *cli.Context) error {
	return mig.Migrate("postgres", c.String("database-dsn"), "./migrations")
}

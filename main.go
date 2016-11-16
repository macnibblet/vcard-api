package main

import (
	"github.com/macnibblet/vcard-api/cmd"
	"gopkg.in/urfave/cli.v2"
	"os"
)

func main() {
	app := cli.App{
		Name:  "VCard-API",
		Usage: "Supported commands http and migrate",
		Commands: []*cli.Command{
			cmd.CMD_HTTP,
			cmd.CMD_MIGRATE,
		},
	}

	app.Run(os.Args)
}

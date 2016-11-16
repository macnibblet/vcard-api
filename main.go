package main

import (
	"bitbucket.org/llg/vcard"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"gopkg.in/urfave/cli.v2"
	"github.com/macnibblet/vcard-api/cmd"
)

func main() {
	app := cli.App{
		Name: "VCard-API",
		Usage: "Supported commands http and migrate",
		Commands: []cli.Command{
			cmd.CMD_HTTP,
			cmd.CMD_MIGRATE,
		},
	}

	app.Run(os.Args)
}

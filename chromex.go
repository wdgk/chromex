package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "chromex"
	app.Version = Version
	app.Usage = ""
	app.Author = "gaku"
	app.Email = "kuga729@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}

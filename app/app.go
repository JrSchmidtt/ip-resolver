package app

import "github.com/urfave/cli"

//Build comand-line app
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Comand line app in go"
	app.Usage = "IP Resolver"
	return app
}
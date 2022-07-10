package app

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	"net"
	)


//Build comand-line app
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Comand line app in go"
	app.Usage = "IP Resolver"
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Lookup ip address of a website.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: findIP,
		},
	}
	return app
}

func findIP(c *cli.Context){
	host := c.String("host")
	fmt.Println("IP's:",host)
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips{
		fmt.Println(ip)
	}
}
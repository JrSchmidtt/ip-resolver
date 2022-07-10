package app

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
	"net")

//Build comand-line app
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Comand line app in go"
	app.Usage = "IP Resolver"
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Lookup ip address of a website.",
			Flags: flags,
			Action: findIP,
		},
		{
			Name: "servers",
			Usage: "Find DNS of host",
			Flags: flags,
			Action: findServers,
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

func findServers(c *cli.Context){
	host := c.String("host")
	fmt.Println("DNS for:",host)
	servers, err := net.LookupNS(host) //name servers
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
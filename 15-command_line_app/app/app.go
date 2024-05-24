package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Web Monitor"
	app.Usage = "Search IPs and server names from internet"
	app.Author = "Marcos H Braz"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.pudim.com.br",
			Usage: "A hostname.",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ips from a host",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Get the name of the server",
			Flags:  flags,
			Action: getServerName,
		},
	}
	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServerName(c *cli.Context) {
	host := c.String("host")
	serverName, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ns := range serverName {
		fmt.Println(ns.Host)
	}
}

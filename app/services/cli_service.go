package services

import (
	"crawler/app/config/routes"

	"github.com/urfave/cli"
)

/*Function with generate mod cli application*/
func GenerateCli() *cli.App {
	app := cli.NewApp()
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "http://quotes.toscrape.com/",
			Usage: "criando um webcrawler usando golang",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "3000",
			Usage: "subir servidor web",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "crawler",
			Usage:  "criando um webcrawler usando golang",
			Flags:  flags,
			Action: call,
		},
		{
			Name:   "server",
			Usage:  "subindo server",
			Flags:  flags,
			Action: server,
		},
	}
	return app
}

func call(c *cli.Context) {
	GenerateCrawler(c.String("host"))
}

func server(c *cli.Context) {
	routes.GenerateWebServer(c.String("port"))
}

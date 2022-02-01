package services

import (
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
	}

	app.Commands = []cli.Command{
		{
			Name:   "crawler",
			Usage:  "criando um webcrawler usando golang",
			Flags:  flags,
			Action: call,
		},
	}
	return app
}

func call(c *cli.Context) {
	GenerateCrawler(c.String("host"))
}

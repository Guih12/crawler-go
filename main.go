package main

import (
	"crawler/app/services"
	"os"
)

func main() {
	app := services.GenerateCli()
	app.Run(os.Args)
}

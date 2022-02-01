package main

import (
	"crawler/app/repositories"
	"crawler/app/services"
	"crawler/app/utils"
	"fmt"
	"os"
)

func main() {
	app := services.GenerateCli()
	app.Run(os.Args)

	ar := repositories.GetAllArticles()
	for _, a := range ar {
		json, _ := utils.GenerateJson(a)
		fmt.Println(string(json))
	}
}

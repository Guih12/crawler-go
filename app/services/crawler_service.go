package services

import (
	"crawler/app/entity"
	"crawler/app/repositories"
	"fmt"

	"github.com/gocolly/colly/v2"
)

func GenerateCrawler(host string) string {
	collection := colly.NewCollector()

	collection.OnHTML(".quote", func(e *colly.HTMLElement) {
		articles := make([]entity.Article, 0)
		newPhrase := e.ChildText(".text")
		newAuthor := e.ChildText(".author")

		article := &entity.Article{
			Phrase: newPhrase,
			Author: newAuthor,
		}

		articles = append(articles, *article)
		articlesRepository := &repositories.ArticlesRepository{Articles: articles}
		articlesRepository.Store()
	})

	collection.OnHTML(".next a", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		collection.Visit(h.Request.AbsoluteURL(link))
	})

	collection.OnRequest(func(response *colly.Request) {
		fmt.Println("Acessou: ", response.URL)
	})

	collection.Visit(host)
	return "Crawler gerado com sucesso"
}

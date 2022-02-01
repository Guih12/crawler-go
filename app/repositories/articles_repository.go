package repositories

import (
	"crawler/app/entity"
	"encoding/json"
	"log"
)

type ArticlesRepository struct {
	articles []entity.Article
}

func NewArticleRepository(articles []entity.Article) *ArticlesRepository {
	return &ArticlesRepository{articles: articles}
}

func (articlesRepository *ArticlesRepository) Store() {

}

func converJson(article entity.Article) []byte {
	json, err := json.Marshal(article)

	if err != nil {
		log.Fatal("Error")
	}
	return json
}

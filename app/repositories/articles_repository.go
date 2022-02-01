package repositories

import (
	"crawler/app/entity"
)

type ArticlesRepository struct {
	articles []entity.Article
}

func NewArticleRepository(articles []entity.Article) *ArticlesRepository {
	return &ArticlesRepository{articles: articles}
}

func (articlesRepository *ArticlesRepository) Store() bool {
	return true
}

package repositories

import (
	"crawler/app/database"
	"crawler/app/entity"
	"crawler/app/observer"
	"log"
)

type ArticlesRepository struct {
	Articles []entity.Article
}

func (articlesRepository ArticlesRepository) Store() {
	for _, article := range articlesRepository.Articles {
		response := persist(&article)
		if response {
			observer := observer.IObserver{Article: article}
			observer.ObserverArticle()
		}
	}
}

func GetAllArticles() []entity.Article {
	database, _ := database.Connect()

	rows, erro := database.Query("SELECT * FROM articles")

	articles := make([]entity.Article, 0)
	for rows.Next() {
		var article entity.Article
		if erro = rows.Scan(&article.ID, &article.Phrase, &article.Author); erro != nil {
			log.Fatal(erro)
		}
		articles = append(articles, article)
	}

	return articles
}

func persist(article *entity.Article) bool {
	db, _ := database.Connect()
	stent, _ := db.Prepare("INSERT INTO articles (phrase, author) values(?, ?)")

	defer stent.Close()

	insert, _ := stent.Exec(*&article.Phrase, *&article.Author)

	if insert != nil {
		return true
	}

	return false
}

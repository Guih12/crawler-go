package observer

import (
	"crawler/app/entity"
	"log"
)

type IObserver struct {
	Article entity.Article
}

func (observer *IObserver) ObserverArticle() bool {
	if observer != nil {
		log.Printf("O artigo :%s do autor: %s foi inserido", *&observer.Article.Phrase, *&observer.Article.Author)
	}

	return false
}

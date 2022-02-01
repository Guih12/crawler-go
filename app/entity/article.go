package entity

type Article struct {
	Phrase string `json: "frase"`
	Author string `json: "autor"`
}

func NewArticle() *Article {
	return &Article{}
}

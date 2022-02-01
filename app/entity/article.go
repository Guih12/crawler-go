package entity

type Article struct {
	ID     int    `json: "id"`
	Phrase string `json: "frase"`
	Author string `json: "autor"`
}

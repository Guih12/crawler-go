package tests

import (
	"crawler/app/entity"
	"crawler/app/utils"
	"testing"
)

func TestGenerateJson(t *testing.T) {
	article := entity.Article{
		Phrase: "Código sem testes não é limpo",
		Author: "David Thomas",
	}

	t.Run("failed json", func(t *testing.T) {
		_, err := utils.GenerateJson(article)

		if err != nil {
			t.Fail()
		}
	})

	t.Run("test return json convert", func(t *testing.T) {
		json, _ := utils.GenerateJson(article)

		if json == nil {
			t.Fail()
		}
	})
}

package tests

import (
	"crawler/app/services"
	"testing"
)

func TestCrawlerService(t *testing.T) {
	generateCralwer := services.GenerateCrawler("http://quotes.toscrape.com/")

	messageExpected := "Crawler gerado com sucesso"

	if generateCralwer != messageExpected {
		t.Failed()
	}
}

package controllers

import (
	"crawler/app/repositories"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	articles := repositories.GetAllArticles()

	if len(articles) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Articles not found"))
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(articles); erro != nil {
		return
	}
}

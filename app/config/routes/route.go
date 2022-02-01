package routes

import (
	"crawler/app/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GenerateWebServer(port string) {
	log.Printf("Server on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), generateRouter()))
}

func generateRouter() *mux.Router {
	router := mux.NewRouter()
	pathRoutes(router)
	return router
}

func pathRoutes(router *mux.Router) {
	router.HandleFunc("/articles", controllers.Index).Methods(http.MethodGet)
}

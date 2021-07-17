package main

import (
	"log"
	"net/http"

	"github.com/kohkb/smash_sp/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// ルート(エンドポイント)
	r.HandleFunc("/videos/search", handler.GetVideos).Methods("GET")
	r.HandleFunc("/videos/favorites", handler.GetVideos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

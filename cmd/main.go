package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/podikoglou/gifthing/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/gif", handlers.GifHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}

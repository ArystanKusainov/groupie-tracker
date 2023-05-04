package main

import (
	"log"
	"net/http"

	"github.com/akusaino/groupie-tracker/cmd/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist/", handlers.ArtistPage)
	http.HandleFunc("/search/", handlers.Search)
	http.HandleFunc("/filters/", handlers.Filters)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	log.Println("Запуск сервера на http://127.0.0.1:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

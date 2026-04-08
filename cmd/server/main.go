package main

import (
	"log"
	"net/http"

	"github.com/halxdocs/ghostapi/internal/api"
)

func main() {
	handler := api.NewHandler()

	http.HandleFunc("/scrape", handler.Scrape)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
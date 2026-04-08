package main

import (
	"log"
	"os"

	"github.com/halxdocs/ghostapi/internal/engine"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/server/main.go <url>")
	}

	url := os.Args[1]

	e := engine.NewEngine()

	if err := e.Run(url); err != nil {
		log.Fatal(err)
	}
}
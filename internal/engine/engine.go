package engine

import (
	"fmt"

	"github.com/halxdocs/ghostapi/internal/scraper"
)

type Engine struct {
	scraper *scraper.Scraper
}

func NewEngine() *Engine {
	return &Engine{
		scraper: scraper.NewScraper(),
	}
}

func (e *Engine) Run(url string) error {
	html, err := e.scraper.FetchHTML(url)
	if err != nil {
		return err
	}

	fmt.Println("HTML fetched successfully")
	fmt.Println("Length:", len(html))

	return nil
}
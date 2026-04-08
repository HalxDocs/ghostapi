package engine

import (
	"encoding/json"
	"fmt"

	"github.com/halxdocs/ghostapi/internal/parser"
	"github.com/halxdocs/ghostapi/internal/scraper"
)

type Engine struct {
	scraper *scraper.Scraper
	parser  *parser.Parser
}

func NewEngine() *Engine {
	return &Engine{
		scraper: scraper.NewScraper(),
		parser:  parser.NewParser(),
	}
}

func (e *Engine) Run(url string) error {
	html, err := e.scraper.FetchHTML(url)
	if err != nil {
		return err
	}

	data, err := e.parser.ParseHTML(html)
	if err != nil {
		return err
	}

	jsonOutput, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(jsonOutput))

	return nil
}
package engine

import (
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

func (e *Engine) Process(url string) (*parser.ParsedData, error) {
	html, err := e.scraper.FetchHTML(url)
	if err != nil {
		return nil, err
	}

	return e.parser.ParseHTML(html)
}
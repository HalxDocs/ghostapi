package scraper

import (
	"github.com/gocolly/colly/v2"
)

type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) FetchHTML(url string) (string, error) {
	c := colly.NewCollector()

	var htmlContent string

	c.OnResponse(func(r *colly.Response) {
		htmlContent = string(r.Body)
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}
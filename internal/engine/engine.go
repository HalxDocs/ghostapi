package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedData struct {
	Title string   `json:"title"`
	Text  []string `json:"text"`
	Links []map[string]string `json:"links"`
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseHTML(html string) (*ParsedData, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	data := &ParsedData{}

	// Title
	data.Title = doc.Find("title").First().Text()

	// Extract ALL visible text (div, span, p)
	doc.Find("body *").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		if text != "" && len(text) > 30 {
			data.Text = append(data.Text, text)
		}
	})

	// Extract links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		text := strings.TrimSpace(s.Text())

		if exists && text != "" {
			data.Links = append(data.Links, map[string]string{
				"text": text,
				"href": href,
			})
		}
	})

	return data, nil
}
package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// =====================
// DATA STRUCTURE
// =====================

type ParsedData struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// =====================
// PARSER
// =====================

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseHTML(html string) (*ParsedData, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	// Detect page type
	if isProductPage(doc) {
		return parseProduct(doc), nil
	}

	if isBlogPage(doc) {
		return parseBlog(doc), nil
	}

	return parseGeneric(doc), nil
}

// =====================
// DETECTION LOGIC
// =====================

func isProductPage(doc *goquery.Document) bool {
	text := strings.ToLower(doc.Text())

	return strings.Contains(text, "add to cart") ||
		strings.Contains(text, "buy now") ||
		strings.Contains(text, "price")
}

func isBlogPage(doc *goquery.Document) bool {
	return doc.Find("article").Length() > 0 ||
		doc.Find("time").Length() > 0
}

// =====================
// PRODUCT PARSER
// =====================

func parseProduct(doc *goquery.Document) *ParsedData {
	data := make(map[string]interface{})

	// Product name
	name := strings.TrimSpace(doc.Find("h1").First().Text())
	if name != "" {
		data["name"] = name
	}

	// Price detection
	price := strings.TrimSpace(doc.Find("[class*=price]").First().Text())
	if price != "" {
		data["price"] = price
	}

	// Images
	var images []string
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists && src != "" {
			images = append(images, src)
		}
	})

	if len(images) > 0 {
		data["images"] = images
	}

	return &ParsedData{
		Type: "product",
		Data: data,
	}
}

// =====================
// BLOG PARSER
// =====================

func parseBlog(doc *goquery.Document) *ParsedData {
	data := make(map[string]interface{})

	// Title
	title := strings.TrimSpace(doc.Find("h1").First().Text())
	if title != "" {
		data["title"] = title
	}

	// Content
	var content []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" && len(text) > 20 {
			content = append(content, text)
		}
	})

	if len(content) > 0 {
		data["content"] = content
	}

	return &ParsedData{
		Type: "blog",
		Data: data,
	}
}

// =====================
// GENERIC PARSER (fallback)
// =====================

func parseGeneric(doc *goquery.Document) *ParsedData {
	data := make(map[string]interface{})

	// Title
	title := strings.TrimSpace(doc.Find("title").Text())
	if title != "" {
		data["title"] = title
	}

	// Extract meaningful text
	var textBlocks []string
	doc.Find("body *").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		if text != "" && len(text) > 40 {
			textBlocks = append(textBlocks, text)
		}
	})

	if len(textBlocks) > 0 {
		data["text"] = textBlocks
	}

	return &ParsedData{
		Type: "generic",
		Data: data,
	}
}
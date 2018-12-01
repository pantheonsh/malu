package utils

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func GoogleSearch(query string) ([]string, error) {
	var protocol = "https://"
	var host = "www.google.com"
	var lang = "pt"
	var links = []string{}

	res, err := http.Get(protocol + host + "/search?q=" + EncodeURIComponent(query) + "&hl=" + EncodeURIComponent(lang))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("Não foi HTTP 200 ao pesquisar no Google.")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("h3.r > a").Each(func(index int, sel *goquery.Selection) {
		href, exists := sel.Attr("href")

		if exists {
			parsed, _ := url.ParseQuery(href)
			var link = parsed.Get("/url?q")
			if link != "" {
				links = append(links, link)
			}
		}
	})

	return links, nil
}

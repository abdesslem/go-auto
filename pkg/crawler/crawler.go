package crawler

import (
	"fmt"
	"go-auto/website"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var searchURL string = "https://www.leboncoin.fr/recherche?category=2&text=bmw"
var nextPage string = "https://www.leboncoin.fr/recherche?category=2&text=bmw&page=2"

func Search(ws *website.Website) string {

	sURL := ws.SearchURL
	fmt.Println(sURL)
	// get next page

}

func Crawl(ws *website.Website, cssSelector string) ([]string, error) {

	URL := ws.URL
	cssSelectorSample := ".post-title"
	cssSelector = cssSelectorSample

	resp, err := http.Get(URL)

	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []string{}, err
	}

	// Save each .post-title as a list
	links := []string{}

	doc.Find(cssSelector).Each(func(i int, s *goquery.Selection) {
		links = append(links, s.Text())
	})

	return links, nil
}

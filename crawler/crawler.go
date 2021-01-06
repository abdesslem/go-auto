package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type WebPage struct {
	name      string
	url       string
	searchURL string
	useProxy  bool
}

type Leboncoin struct {
	webpage WebPage
	another string
}

type Lacentrale WebPage

type Scrapper interface {
	scrape()
}

func (wp *WebPage) scrape() {
	fmt.Print("Go scrape")
}

func (lbc *Leboncoin) scrape() {
	fmt.Print("Go scrape")
}

var aweb WebPage = WebPage{
	"leboncoin",
	"https://leboncoin.fr",
	"searchhttp",
	true,
}

func Justatestcase() {
	tt := Leboncoin{
		aweb,
		"another",
	}
	fmt.Print(tt)
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestBlogTitles(url string) (string, error) {

	// Get the HTML
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Save each .post-title as a list
	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "- " + s.Text() + "\n"
	})
	return titles, nil
}

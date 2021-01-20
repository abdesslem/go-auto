package scrapper

import (
	"go-auto/website"

	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(html *goquery.Document) ([]website.Car, error) {

	var crawlerCar = []string{}
	brandSelection := {}
}

func GetBrand(html *goquery.Document, cssSelector string) string {

	html.Find(cssSelector).Each(func(i int, s *goquery.Selection) {
		return s.Text()[0]
	})
}

func GetBrand(html *goquery.Document, cssSelector string) string {

	html.Find(cssSelector).Each(func(i int, s *goquery.Selection) {
		return s.Text()[0]
	})
}
package main

import (
	"fmt"
	"go-auto/crawler"
	"log"
)

func main() {
	blogTitles, err := crawler.GetLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles:")
	fmt.Printf(blogTitles)

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExampleScrape parses a page
func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://sdcc.sourceforge.net/snap.php#Linux")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("body div table tbody tr td div table tbody tr td").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, ok := s.Attr("href")
			if ok {
				title := s.Text()
				if strings.Contains(title, "sdcc-snapshot-amd64") && strings.Contains(title, "tar.bz2") {
					str := strings.Split(link, "/download")
					fmt.Println(str[0])
					os.Exit(0)
				}
			}
		})
	})
}

func main() {
	ExampleScrape()
}

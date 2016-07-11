package jobs

import (
    "fmt"
	"github.com/PuerkitoBio/goquery"
)

func Crawler(link string) {

	fmt.Println("Crawler Start...")

	doc, _ := goquery.NewDocument(link)

	doc.Find(".r").Each(func(i int, s *goquery.Selection) {
		band , _ := s.Find("a").Attr("href")
		fmt.Println(band)
	})


}





package jobs

import (
	"github.com/revel/revel"
	"github.com/revel/modules/jobs/app/jobs"
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/janekolszak/revmgo"
	"hafiyye/app/models"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"gopkg.in/mgo.v2/bson"
	"net/url"
)


var db *mgo.Database

type KeywordCrawler struct{
}

func (c KeywordCrawler) Run() {
	fmt.Println("Keyword crawler started")
	FetchUrl()

}


func FetchUrl(){
	result := [] models.Keywords{}
	db.C("Keywords").Find(nil).All(&result)

	google := "https://www.google.com/search?tbs=lr:lang_1tr&lr=lang_tr&q="
	for _ , r := range  result {
		Crawler( google + url.QueryEscape(r.Keyword), r.Keyword)
	}
}


func Crawler(link, keyword string) {

	fmt.Println("Crawler Start..." + link)

	doc, _ := goquery.NewDocument(link)

	banlist :=[] models.Banlist{}
	db.C("Banlist").Find(nil).All(&banlist)

	doc.Find(".r").Each(func(i int, s *goquery.Selection) {
		band , _ := s.Find("a").Attr("href")

		link := strings.Index(band, "&sa")
		href := band[7:link]

		if strings.Index(href,"http") > -1 {

			InsertDb(href, keyword)


		}
	})

}

func InsertDb(url, keyword string){

	data := models.Document{}
	db.C("Document").Find(bson.M{"Url":url}).One(&data)

	if len(data.Url) < 2 {

		doc := models.Document{}
		doc.Keyword = keyword
		doc.Url = url

		fmt.Println(url)

		//db.C("Document").Insert(&doc)

	}
}



func init() {

	revel.OnAppStart(func() {

		db = revmgo.Session.DB(revel.Config.StringDefault("mongo.database", "local"))
		jobs.Schedule("@every 1m", KeywordCrawler{})
	})
}
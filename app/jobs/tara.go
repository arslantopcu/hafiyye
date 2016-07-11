package jobs

import (
	"github.com/revel/revel"
	"github.com/revel/modules/jobs/app/jobs"
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/janekolszak/revmgo"
	"hafiyye/app/models"
)


var db *mgo.Database

type KeywordCrawler struct{
}

func (c KeywordCrawler) Run() {
	fmt.Println("Keyword crawler started")

	result := [] models.Keywords{}

	db.C("Keywords").Find(nil).Sort("-timestamp").All(&result)

	google := "https://www.google.com/search?q="
	for _ , r := range  result {

		fmt.Println(google + r.Keyword)
		//Crawler(google + r.Keyword)

	}

}


func init() {

	revel.OnAppStart(func() {

		db = revmgo.Session.DB(revel.Config.StringDefault("mongo.database", "local"))
		jobs.Schedule("@every 1m", KeywordCrawler{})
	})
}
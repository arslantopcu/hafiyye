package controllers

import (
	"github.com/janekolszak/revmgo"
	"gopkg.in/mgo.v2"
	"github.com/revel/revel"
)

var db *mgo.Database

func InitMongoDb(){

	revmgo.AppInit()

	revmgo.Session.SetPoolLimit( revel.Config.IntDefault("mongo.maxPool", 10))

	// Database ve Documan ayarlanıyor...
	db = revmgo.Session.DB(revel.Config.StringDefault("mongo.database", "local"))

}
package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Document struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name 	string
	Keyword	string
	Url 	string
	Timestamp time.Time
}

type Banlist struct {
	Id 	bson.ObjectId	`bson:"_id,omitempty"`
	Url 	string
}


type Keywords struct {
	Id 		bson.ObjectId	`bson:"_id,omitempty"`
	Name		string
	Keyword 	string
	Timestamp	time.Time
}

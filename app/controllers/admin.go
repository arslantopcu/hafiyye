package controllers

import (
	"github.com/revel/revel"
	"hafiyye/app/routes"
	"gopkg.in/mgo.v2/bson"
	"hafiyye/app/models"
	"fmt"
	"time"
	"math"
)

type Admin struct {
	*revel.Controller
}


func (c Admin) checkUser() revel.Result {

	if c.Session["username"] != "aslan" {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	return nil
}


func (c Admin) Linkler(p int) revel.Result {

	if p == 0 {
		p=1
	}

	sizeLimit := 20
	page := (p-1)*sizeLimit
	result := [] models.Document{}

	db.C("Document").Find(nil).Sort("-timestamp").Skip(page).Limit(sizeLimit).All(&result)

	t, err:= db.C("Document").Find(nil).Count()

	if err != nil{
		fmt.Println("Sayfa yok")
	}

	username := c.Session["username"]

	next := p+1
	prev := p-1
	total := math.Ceil(float64(t)/float64(sizeLimit))

	return c.Render(username, result, total, p, next, prev)
}


func (c Admin) Banlist(p int) revel.Result {

	if p == 0 {
		p=1
	}

	sizeLimit := 20
	page := (p-1)*sizeLimit
	result := [] models.Banlist{}

	db.C("Banlist").Find(nil).Sort("-timestamp").Skip(page).Limit(sizeLimit).All(&result)

	t, err:= db.C("Banlist").Find(nil).Count()

	if err != nil{
		fmt.Println("Sayfa yok")
	}
	username := c.Session["username"]

	next := p+1
	prev := p-1
	total := math.Ceil(float64(t)/float64(sizeLimit))

	return c.Render(username, result, total, p, next, prev)
}

func (c Admin) Banlistekle() revel.Result {
	return c.Render()
}

func (c Admin) Banlistinsert(url string) revel.Result {

	data := models.Banlist{}
	db.C("Banlist").Find(bson.M{"Url": url}).One(&data)

	if len(data.Url) > 2 {
		c.Flash.Error("Url adresi kayitli")
		return c.Redirect(routes.Admin.Banlistekle())
	}

	key := models.Banlist{}
	key.Url = url

	db.C("Banlist").Insert(&key)

	return c.Redirect(routes.Admin.Banlist(0))
}

func (c Admin) BanlistDelete(id string) revel.Result {

	db.C("Banlist").RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})

	return c.Redirect(routes.Admin.Banlist(0))
}

func (c Admin) LinklerDelete(id string) revel.Result {

	db.C("Document").RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})

	return c.Redirect(routes.Admin.Linkler(0))
}



func (c Admin) Index(p int) revel.Result {

	if p == 0 {
		p=1
	}

	sizeLimit := 20
	page := (p-1)*sizeLimit
	result := [] models.Keywords{}

	db.C("Keywords").Find(nil).Sort("-timestamp").Skip(page).Limit(sizeLimit).All(&result)

	t, err:= db.C("Keywords").Find(nil).Count()

	if err != nil{
		fmt.Println("Sayfa yok")
	}

	username := c.Session["username"]

	next := p+1
	prev := p-1
	total := math.Ceil(float64(t)/float64(sizeLimit))


	return c.Render(username, result, total, p, next, prev)
}

func (c Admin) InsertDocument() revel.Result {
	return c.Render()
}

func (c Admin) DeleteDocument(id string) revel.Result {

	db.C("Keywords").RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})

	return c.Redirect(routes.Admin.Index(0))
}

func (c Admin) SaveKeyword(name, keyword string) revel.Result {

	key := models.Keywords{}
	key.Name = name
	key.Keyword = keyword
	key.Timestamp = time.Now()

	data := models.Keywords{}
	db.C("Keywords").Find(bson.M{"Name": keyword}).One(&data)

	if len(data.Name) > 2 {
		c.Flash.Error("Keyword adresi kayitli")
		return c.Redirect(routes.Admin.InsertDocument())
	}

	db.C("Keywords").Insert(&key)

	return c.Redirect(routes.Admin.Index(0))
}

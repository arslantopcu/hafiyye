package controllers

import (
	"github.com/revel/revel"
	"hafiyye/app/routes"
	"github.com/janekolszak/revmgo"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (c App) Index() revel.Result {
	return c.Render()
}


func (c App) Login() revel.Result {
	return c.Render()
}


func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) LoginCkeck(username, password string) revel.Result {

	if username == "admin" && password == "12345" {

		c.Session["username"]="admin"

		return  c.Redirect(Admin.Index)
	}else{
		c.Flash.Error(c.Message("Login hatasi"))
		return c.Redirect(App.Login)
	}

}
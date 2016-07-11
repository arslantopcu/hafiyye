package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitMongoDb)
	revel.InterceptMethod(Admin.checkUser, revel.BEFORE)

}

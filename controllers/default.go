package controllers

import (
	"github.com/astaxie/beego"
)

// MainController main controller
type MainController struct {
	beego.Controller
}

// Get get way
func (c *MainController) Get() {
	c.TplName = "index.html"
}

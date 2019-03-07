package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// LoginController login
type LoginController struct {
	beego.Controller
}

// Post ....
func (l *LoginController) Post() {
	if l.Ctx.Input.Param("username") == "1" {
		l.Ctx.WriteString("Hello World")
	}
	l.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	l.Redirect("/", 302)
}

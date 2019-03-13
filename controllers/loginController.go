package controllers

import (
	"net/http"

	"TyWeb/models/dao"

	"github.com/astaxie/beego"
)

// LoginController login
type LoginController struct {
	beego.Controller
}

// Post ....
func (l *LoginController) Post() {

	name := l.GetString("username")
	passwd := l.GetString("passwd")
	ok := dao.SelectUserTable(name, passwd)
	if ok {
		l.Ctx.WriteString("Hello World")
	}
	l.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	l.Data["json"] = "用户名或密码错误"
	l.ServeJSON()
	l.Finish()
}

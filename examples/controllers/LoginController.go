package controllers

import (
	"github.com/danuxguin/daxweb"
	"net/http"
)

type LoginController struct {
	daxweb.Controller
}

func (this *LoginController) Handler(w http.ResponseWriter, r *http.Request) {
	this.Controller.Init(this, w, r)
}

func (this *LoginController) Get() {
	this.Tpl = "./views/login.html"
	this.ExecuteTpl()
}

func (this *LoginController) Post() {
	this.Ctx.R.ParseForm()
	account := this.Ctx.R.FormValue("account")
	password := this.Ctx.R.FormValue("password")

	this.SetCookie("account", account, 5*60)
	this.SetCookie("password", password, 5*60)

	if account == "admin" && password == "admin" {
		http.Redirect(this.Ctx.W, this.Ctx.R, "/", http.StatusFound)
	}
}

package controllers

import (
	"fmt"
	"github.com/danuxguin/daxweb"
	"net/http"
)

type MainController struct {
	daxweb.Controller
}

func (this *MainController) Handler(w http.ResponseWriter, r *http.Request) {
	this.Controller.Init(this, w, r)
}

func (this *MainController) Get() {

	account := this.GetCookie("account")
	password := this.GetCookie("password")

	if account == "admin" && password == "admin" {
		this.Tpl = "./views/database.html"
		this.ExecuteTpl()
	} else {
		http.Redirect(this.Ctx.W, this.Ctx.R, "/login", http.StatusFound)
	}

}

func (this *MainController) Post() {
	this.Ctx.R.ParseForm()
	ip := this.Ctx.R.FormValue("ip")
	port := this.Ctx.R.FormValue("port")
	group := this.Ctx.R.FormValue("group")
	db := this.Ctx.R.FormValue("db")

	fmt.Printf("ip=%v port=%v group=%v db=%v\n", ip, port, group, db)

	fmt.Fprintf(this.Ctx.W, "Add database sucessfule!")
}

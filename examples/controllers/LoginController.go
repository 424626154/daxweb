package controllers

import (
	"fmt"
	"github.com/danuxguin/daxweb"
	"net/http"
)

type LoginController struct {
	daxweb.Controller
}

func (this *LoginController) Handler(w http.ResponseWriter, r *http.Request) {
	this.Controller.Init(this, w, r)
}

func (this *LoginController) Login() {
	if this.R.Method == "GET" {
		this.Tpl = "./views/login.html"
		this.ExecuteTpl()
	} else if this.R.Method == "POST" {
		fmt.Fprintf(this.W, "Login----------------------Post")
	}
}

func (this *LoginController) Get() {
	fmt.Fprintf(this.W, "----------------------Get")
}

func (this *LoginController) Post() {
	fmt.Fprintf(this.W, "----------------------Post")
}

package main

import (
	"github.com/danuxguin/daxweb"
	"webmanager/controllers"
)

func init() {
	daxweb.Dax.AddRouter("/", &controllers.MainController{})
	daxweb.Dax.AddRouter("/login", &controllers.LoginController{})
}

func main() {
	daxweb.Run("127.0.0.1", 8888)
}

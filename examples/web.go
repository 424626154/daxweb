package main

import (
	"examples/controllers"
	"fmt"
	"github.com/danuxguin/daxweb"
)

func init() {
	daxweb.Dax.AddRouter("/login", &controllers.LoginController{})
}

func main() {
	daxweb.Run("127.0.0.1", 9999)
}

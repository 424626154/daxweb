package daxweb

import (
	"fmt"
	"github.com/danuxguin/daxweb/common"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

type Controller struct {
	Data map[interface{}]interface{}
	Tpl  string

	Ctx
}

func (this *Controller) Init(c interface{}, w http.ResponseWriter, r *http.Request) {
	this.Data = make(map[interface{}]interface{})

	this.Ctx.W = w
	this.Ctx.R = r

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %v\n", err)
			http.NotFound(w, r)
		}
	}()

	urlList := common.GetUrlSplit(r.URL.Path)
	control := reflect.ValueOf(c)
	if len(urlList) == 1 {
		if r.Method == "GET" {
			control.MethodByName("Get").Call(nil)
		} else if r.Method == "POST" {
			control.MethodByName("Post").Call(nil)
		}
	} else {
		preurl := urlList[1]
		urlFunc := strings.ToUpper(preurl[:1]) + preurl[1:]
		control.MethodByName(urlFunc).Call(nil)
	}
}

func (this *Controller) ExecuteTpl() {
	t, err := template.ParseFiles(this.Tpl)
	if err != nil {
		http.NotFound(this.Ctx.W, this.Ctx.R)
		return
	}

	t.Execute(this.Ctx.W, this.Data)
}

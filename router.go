package daxweb

import (
	"github.com/danuxguin/daxweb/common"
	"net/http"
	"sync"
)

var Dax dax

type dax struct {
	sync.RWMutex
}

type Router interface {
	Handler(http.ResponseWriter, *http.Request)
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func init() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
}

func (this *dax) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := common.GetHeader(r.URL.Path)
	h, ok := mux[head]
	if ok {
		h(w, r)
		return
	}

	http.NotFound(w, r)
}

func (this *dax) AddRouter(path string, r Router) {
	this.Lock()
	defer this.Unlock()

	mux[path] = r.Handler
}

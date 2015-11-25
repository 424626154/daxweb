package daxweb

import (
	"fmt"
	"net/http"
	"time"
)

type Ctx struct {
	W http.ResponseWriter
	R *http.Request
}

func (this *Ctx) SetCookie(key string, value string, expires int64) {

	var cookie http.Cookie
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Unix(time.Now().Unix()+expires, 0)

	this.R.AddCookie(&cookie)

	this.W.Header().Add("Set-Cookie", cookie.String())
}

func (this *Ctx) GetCookie(key string) string {
	cookie, err := this.R.Cookie(key)
	if err == http.ErrNoCookie {
		fmt.Printf("%v ERR = ErrNoCookie\n", key)
		return ""
	}

	return cookie.Value
}

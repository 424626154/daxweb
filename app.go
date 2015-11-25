package daxweb

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Session *http.Server

func Run(ip string, port int) {
	Session = &http.Server{
		Addr:           fmt.Sprintf("%v:%v", ip, port),
		Handler:        &Dax,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(Session.ListenAndServe())
}

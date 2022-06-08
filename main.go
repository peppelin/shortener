package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	content = "Content-type"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte("<h1>Hello world</h1>"))
	})

	fmt.Printf("starting server at 127.0.0.1%s\n", s.Addr)
	log.Fatal(s.ListenAndServe())

}

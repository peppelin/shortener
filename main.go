package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	content = "Content-type"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func main() {

	port := ":8080"

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte("<h1>Hello world</h1>"))
	})

	fmt.Println("serving ...")
	http.ListenAndServe(port, r)

}

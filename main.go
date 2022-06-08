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

	r.Get("/api/v1/lookup/{urlCode}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte(chi.URLParam(r, "urlCode")))
	})

	r.Get("/api/v1/shorten/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte(chi.URLParam(r, "url")))
	})

	fmt.Println("serving ...")
	http.ListenAndServe(port, r)

}

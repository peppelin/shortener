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

	// mock data for testing

	shortened := map[string]string{
		"11111": "www.google.com",
		"222":   "www.222.com",
	}
	r.Get("/api/v1/lookup/{urlCode}", handleShort(shortened))
	// r.Get("/api/v1/lookup/{urlCode}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set(content, html)
	// 	w.Write([]byte(chi.URLParam(r, "urlCode")))
	// })

	r.Get("/api/v1/shorten/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte(chi.URLParam(r, "url")))
	})

	fmt.Println("serving ...")
	http.ListenAndServe(port, r)

}
func handleShort(shortened map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		short := chi.URLParam(r, "urlCode")
		if long, ok := shortened[short]; ok {
			w.Write([]byte(long))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		}
	}
}

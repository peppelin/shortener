package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"strconv"

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

	r.Get("/api/v1/lookup/{urlCode}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		code := chi.URLParam(r, "urlCode")
		w.Write([]byte(shortened[code]))
	})

	r.Get("/api/v1/shorten/{url}", shorten(shortened))

	fmt.Println("serving ...")
	http.ListenAndServe(port, r)

}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))

	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

func shorten(shortened map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		//w.Write([]byte(chi.URLParam(r, "url")))

		// create a hash from url with a max of 10 character
		hashed := hash(chi.URLParam(r, "url"))

		w.Write([]byte(hashed))

		shortened[hashed] = chi.URLParam(r, "url")
	}
}

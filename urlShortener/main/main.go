package main

import (
	"fmt"
	"net/http"

	"github.com/regiets/gophercises/urlShortener"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-gdoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":    "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlShortener.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final
`
	yamlHandler, err := urlShortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

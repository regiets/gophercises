# URL Shortener

A url shortener that looks at the path of any incomming web request and determines if it should redirect the user to a new page.

### Getting started
1. Start the URL shortener
  ```sh
  go run main/main.go
  ```
2. Try different URLs (below) in your browser

### URLS
- `**http://localhost:8080/:**` uses the defaultMux and returns "hello world"
- `**http://localhost:8080/urlshort-gdoc:**` uses the mapHandler and redirects to `https://godoc.org/github.com/gophercises/urlshort`
- `**http://localhost:8080/yaml-godoc:**` uses the mapHandler and redirects to `https://godoc.org/gopkg.in/yaml.v2`
- `**http://localhost:8080/urlshort:**` uses the yamlHandler and redirects to `https://github.com/gophercises/urlshort`
- `**http://localhost:8080/urlshort-final:**` uses the yamlHandler and redirects to `https://github.com/gophercises/urlshort/tree/final`

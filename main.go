package main

import (
	"fmt"
	"gophercises/url-shorterner/handler"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/test": "https://google.com",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /google
  url: https://google.com
`

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting the serve on :8080")
	http.ListenAndServe(":8080", yamlHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

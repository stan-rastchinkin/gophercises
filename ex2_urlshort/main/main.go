package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sharkstoned/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := loadYamlConfig("./config.yaml")
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")

	http.ListenAndServe(":8080", yamlHandler)
}

func loadYamlConfig(pathToFile string) ([]byte, error) {
	// This is a very basic method that loads the whole file into memory
	// Opening file, reading it line-by-line and closing can make more sense
	// memory-wise in a differet case
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

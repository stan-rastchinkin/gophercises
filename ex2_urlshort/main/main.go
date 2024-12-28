package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sharkstoned/gophercises/urlshort"
	config "github.com/sharkstoned/gophercises/urlshort/config"
)

func main() {
	mux := defaultMux()

	// // Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	configData, err := readFile("./config.yaml")
	panicIfError(err)

	config, err := config.LoadYamlConfig(configData)
	panicIfError(err)

	handler := urlshort.MapHandler(config, mux)

	fmt.Println("Starting the server on :8080")

	http.ListenAndServe(":8080", handler)
}

func readFile(pathToFile string) ([]byte, error) {
	// This is a very basic method that loads the whole file into memory
	// Opening file, reading it line-by-line and closing can make more sense
	// memory-wise in a differet case
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

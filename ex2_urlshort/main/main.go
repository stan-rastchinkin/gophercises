package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/sharkstoned/gophercises/urlshort"
	config "github.com/sharkstoned/gophercises/urlshort/config"
)

func main() {
	var confPath, confFormat string
	flag.StringVar(&confPath, "c", "", "Path to json or yaml config")
	flag.StringVar(&confFormat, "f", "", "Config format (yaml|json)")

	flag.Parse()

	if confPath == "" {
		fmt.Println("No path to config file provided")
		os.Exit(1)
	}
	if confFormat == "" {
		fmt.Println("No format for config is provided")
		os.Exit(1)
	}

	// Put this rules into Bbolt

	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }

	var rules config.Config

	switch confFormat {
	case "yaml":
		configData, err := readFile(confPath)
		panicIfError(err)

		rules, err = config.LoadYamlConfig(configData)
		panicIfError(err)

	case "json":
		configData, err := readFile(confPath)
		panicIfError(err)

		rules, err = config.LoadJsonConfig(configData)
		panicIfError(err)
	default:
		fmt.Printf("Unknown config format %s", confFormat)
		os.Exit(1)
	}

	handler := urlshort.MapHandler(rules, defaultMux())

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

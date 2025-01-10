package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Printf("Failed to fetch page: %e", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %e", err)
		os.Exit(1)
	}

	fmt.Print(string(body))
}

package urlshort

import (
	"net/http"

	config "github.com/sharkstoned/gophercises/urlshort/config"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls config.Config, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectUrl, pathExists := pathsToUrls[r.URL.Path]

		if pathExists {
			http.Redirect(w, r, redirectUrl, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

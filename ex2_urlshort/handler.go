package urlshort

import (
	"net/http"

	yaml "github.com/go-yaml/yaml"
)

// TODO: purpose of these quotes, last column?
type UnmarshalledYamlItem struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectUrl, pathExists := pathsToUrls[r.URL.Path]

		if pathExists {
			http.Redirect(w, r, redirectUrl, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var unmarshalled []UnmarshalledYamlItem
	err := yaml.UnmarshalStrict(yml, &unmarshalled)
	if err != nil {
		return nil, err
	}

	pathMap := make(map[string]string)
	for _, item := range unmarshalled {
		pathMap[item.Path] = item.Url
	}

	return MapHandler(pathMap, fallback), nil
}

// TODO: discover the possibilities of go mod (tidy?)

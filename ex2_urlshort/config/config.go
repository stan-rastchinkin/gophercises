package config

type Config map[string]string

type redirectRule struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url"  json:"url"`
}

func buildPathMap(parsed []redirectRule) Config {
	pathMap := make(map[string]string)
	for _, item := range parsed {
		pathMap[item.Path] = item.Url
	}

	return pathMap
}

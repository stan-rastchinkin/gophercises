package config

import (
	yaml "github.com/go-yaml/yaml"
)

// TODO: purpose of these quotes, last column?
type UnmarshalledYamlItem struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYaml(yml []byte) ([]UnmarshalledYamlItem, error) {
	var unmarshalled []UnmarshalledYamlItem
	err := yaml.UnmarshalStrict(yml, &unmarshalled)
	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func buildPathMap(parsed []UnmarshalledYamlItem) Config {
	pathMap := make(map[string]string)
	for _, item := range parsed {
		pathMap[item.Path] = item.Url
	}

	return pathMap
}

func LoadYamlConfig(data []byte) (Config, error) {
	parsed, err := parseYaml(data)
	if err != nil {
		return nil, err
	}

	return buildPathMap(parsed), nil
}

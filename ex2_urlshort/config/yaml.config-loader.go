package config

import (
	yaml "github.com/go-yaml/yaml"
)

func parseYaml(yml []byte) ([]redirectRule, error) {
	var unmarshalled []redirectRule
	err := yaml.UnmarshalStrict(yml, &unmarshalled)
	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func LoadYamlConfig(data []byte) (Config, error) {
	parsed, err := parseYaml(data)
	if err != nil {
		return nil, err
	}

	return buildPathMap(parsed), nil
}

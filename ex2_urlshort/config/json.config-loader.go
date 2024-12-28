package config

import "encoding/json"

func parseJson(yml []byte) ([]redirectRule, error) {
	var unmarshalled []redirectRule
	err := json.Unmarshal(yml, &unmarshalled)
	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func LoadJsonConfig(data []byte) (Config, error) {
	parsed, err := parseJson(data)
	if err != nil {
		return nil, err
	}

	return buildPathMap(parsed), nil
}

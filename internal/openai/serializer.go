package openai

import "encoding/json"

func UnmarshalConfig(data []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func MarshalConfig(config *Config) ([]byte, error) {
	return json.Marshal(config)
}

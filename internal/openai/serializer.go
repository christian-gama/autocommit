package openai

import "encoding/json"

// UnmarshalConfig unmarshals a Config from a byte slice.
func UnmarshalConfig(data []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// MarshalConfig marshals a Config into a byte slice.
func MarshalConfig(config *Config) ([]byte, error) {
	return json.Marshal(config)
}

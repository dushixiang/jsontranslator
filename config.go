package jsontranslator

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LLM   LargeLanguageModelConfig `json:"llm"`
	Input string                   `json:"input"`
	From  string                   `json:"from"`
	To    []string                 `json:"to"`
}

func ReadConfig(configPath string) (*Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}
	return &config, nil
}

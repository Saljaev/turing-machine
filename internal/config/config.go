package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ZeroCell          string `yaml:"zero_cell"`
	FullCell          string `yaml:"full_cell"`
	DefaultTapeLength int    `yaml:"default_tape_length"`
	DefaultCodeLength int    `yaml:"default_code_length"`
}

func ConfigLoad() *Config {
	yamlFile, err := os.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatal("failed to read config.yaml")
	}

	var cfg Config

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("unmarshal failed with error: %v", err)
	}

	return &cfg
}

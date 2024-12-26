package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Handler struct {
	Name    string   `toml:"name"`
	Pattern string   `toml:"pattern"`
	Command []string `toml:"command"`
}

type Config struct {
    Handlers []Handler `toml:"handlers"`
}

func Load(configPath string) (*Config, error) {
    var cfg Config
    _, err := toml.DecodeFile(configPath, &cfg)
    if err != nil {
        return nil, fmt.Errorf("error loading config: %w", err)
    }
    return &cfg, nil
}

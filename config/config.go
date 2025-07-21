package config

import (
    "github.com/pelletier/go-toml"
    "os"
)

type Config struct {
    Theme            string
    BatteryThreshold int
    AutoSwitch       bool
}

var DefaultConfig = Config{
    Theme:            "nord",
    BatteryThreshold: 20,
    AutoSwitch:       true,
}

func LoadConfig(path string) (Config, error) {
    file, err := os.ReadFile(path)
    if err != nil {
        return DefaultConfig, err
    }
    var cfg Config
    toml.Unmarshal(file, &cfg)
    return cfg, nil
}
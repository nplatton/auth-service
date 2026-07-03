package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	defaultCfg := &Config{
		Server: Server{
			Host: DefaultServerHost,
			Port: DefaultServerPort,
		},
		Logging: Logging{
			Level: DefaultLogLevel,
		},
	}

	cfg := *defaultCfg

	file, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return defaultCfg, nil
	}
	if err != nil {
		return defaultCfg, err
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return defaultCfg, err
	}

	err = ValidateConfig(&cfg)
	if err != nil {
		return defaultCfg, err
	}

	return &cfg, nil
}

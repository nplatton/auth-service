package config

import "fmt"

type Config struct {
	Server  Server  `yaml:"server"`
	Logging Logging `yaml:"logging"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Logging struct {
	Level string `yaml:"level"`
}

func (cfg *Config) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
}

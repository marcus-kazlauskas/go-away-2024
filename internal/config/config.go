package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	yaml "gopkg.in/yaml.v3"
)

var FilePath = "../properties/go-away-2024.yml"
var ServerCfg *ServerConfig
var DatabaseCfg *DatabaseConfig

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl-mode"`
}

func Load() {
	config, err := loadConfig(FilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ServerCfg = &config.Server
	DatabaseCfg = &config.Database
	log.Infof("Loaded properties from %s", FilePath)
}

func loadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling YAML: %v", err)
	}

	return &config, nil
}

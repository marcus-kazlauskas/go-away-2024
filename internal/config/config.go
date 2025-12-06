package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	yaml "gopkg.in/yaml.v3"
)

const FilePath string = "../properties/go-away-2024.yml"

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	S3       S3Config       `yaml:"s3"`
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

type S3Config struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access-key"`
	SecretKey string `yaml:"secret-key"`
	SslMode   bool   `yaml:"ssl-mode"`
	Region    string `yaml:"region"`
}

func GetConfig() *Config {
	config, err := loadConfig(FilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return config
}

func loadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return &config, nil
}

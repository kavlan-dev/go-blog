package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string // dev, prod
	CORS        []string
	DB          databaseConfig
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
}

func InitConfig() (*Config, error) {
	var config Config
	config.Environment = envOrDefault("ENV", "prod")
	config.CORS = strings.Split(envOrDefault("CORS", "*"), ",")
	config.DB.Host = envOrDefault("DB_HOST", "db")
	config.DB.User = envOrDefault("DB_USER", "")
	config.DB.Password = envOrDefault("DB_PASSWORD", "")
	config.DB.Name = envOrDefault("DB_NAME", "")

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Не верно указано окружение %s", config.Environment)
	}

	if config.DB.Host == "" || config.DB.User == "" || config.DB.Password == "" || config.DB.Name == "" {
		return nil, fmt.Errorf("Не верно указаны настройки базы данных")
	}

	return &config, nil
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}

	return value
}

package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string // dev, prod
	JWTSecret   string
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
	config := Config{
		Environment: envOrDefault("ENV", "prod"),
		JWTSecret:   envOrDefault("JWT_SECRET", ""),
		CORS:        strings.Split(envOrDefault("CORS", "*"), ","),
		DB: databaseConfig{
			Host:     envOrDefault("DB_HOST", "localhost"),
			User:     envOrDefault("DB_USER", ""),
			Password: envOrDefault("DB_PASSWORD", ""),
			Name:     envOrDefault("DB_NAME", ""),
		},
	}

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Не верно указано окружение %s", config.Environment)
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("Не указан jwt")
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

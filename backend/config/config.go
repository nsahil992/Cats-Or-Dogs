package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig holds database connection parameters
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Config holds all application configuration
type Config struct {
	Port     string
	DBConfig DBConfig
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Set default port if not specified
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get database configuration
	dbConfig := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   getEnv("DB_NAME", "catsordogs"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Validate essential configuration
	if dbConfig.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is required")
	}

	return &Config{
		Port:     port,
		DBConfig: dbConfig,
	}, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

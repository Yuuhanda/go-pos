package config

import (
    "os"
)

// DBConfig holds database connection parameters
type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

// GetDBConfig returns the database configuration
func GetDBConfig() *DBConfig {
    return &DBConfig{
        Host:     getEnv("DB_HOST", "127.0.0.1"),
        Port:     getEnv("DB_PORT", "3306"),
        User:     getEnv("DB_USER", "username"),
        Password: getEnv("DB_PASSWORD", "password"),
        DBName:   getEnv("DB_NAME", "go-pos"),
    }
}

// Helper function to get environment variables with fallback
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

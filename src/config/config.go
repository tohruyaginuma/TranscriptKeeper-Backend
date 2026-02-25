package config

import (
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type Config struct {
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

func Load() *Config {
	fmt.Println(lib.GetEnv("DB_HOST", "localhost"))
	fmt.Println(lib.GetEnv("DB_PORT", "5432"))
	fmt.Println(lib.GetEnv("DB_USER", "test_local"))
	fmt.Println(lib.GetEnv("DB_PASSWORD", "password"))
	fmt.Println(lib.GetEnv("DB_NAME", "test_local"))
	fmt.Println(lib.GetEnv("DB_SSL_MODE", "disable"))

	return &Config{
		DBHost:    lib.GetEnv("DB_HOST", "localhost"),
		DBPort:    lib.GetEnv("DB_PORT", "5432"),
		DBUser:    lib.GetEnv("DB_USER", "test_local"),
		DBPass:    lib.GetEnv("DB_PASSWORD", "password"),
		DBName:    lib.GetEnv("DB_NAME", "test_local"),
		DBSSLMode: lib.GetEnv("DB_SSL_MODE", "disable"),
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName,
	)
}

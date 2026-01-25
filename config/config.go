package config

import (
	"fmt"
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
	return &Config{
		DBHost:    "localhost",
		DBPort:    "5432",
		DBUser:    "test_local",
		DBPass:    "password",
		DBName:    "test_local",
		DBSSLMode: "disable",
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName,
	)
}

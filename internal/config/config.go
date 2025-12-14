package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	SSLMode    string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:    os.Getenv("DB_SSLMODE"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

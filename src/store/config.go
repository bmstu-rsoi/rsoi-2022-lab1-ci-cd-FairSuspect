package store

import "os"

// Config ...
type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	dsn := os.Getenv("DATABASE_URL")
	if len(dsn) == 0 {
		dsn = "host=localhost dbname=persons user=program password=test port=5432 sslmode=disable"

	}
	return &Config{
		DatabaseURL: dsn,
	}
}

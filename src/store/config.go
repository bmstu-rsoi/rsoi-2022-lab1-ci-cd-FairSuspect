package store

// Config ...
type Config struct {
	DatabaseURL string "host=localhost:5432 dbname=persons user=program password=test"
}

func NewConfig() *Config {
	return &Config{}
}

package apiserver

import "http-rest-api/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bond_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

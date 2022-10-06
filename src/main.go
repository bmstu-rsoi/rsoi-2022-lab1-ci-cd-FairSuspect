package main

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"http-rest-api/store"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
	repo       store.PersonRepository
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	storeConfig := store.NewConfig()
	st := store.New(storeConfig)
	st.Open()
	repo = *st.Person()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

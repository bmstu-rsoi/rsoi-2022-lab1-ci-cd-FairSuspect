package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ..
)

type Store struct {
	config           *Config
	db               *sql.DB
	personRepository *PersonRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {

	log.Default().Println("Connecting to db... with " + s.config.DatabaseURL)

	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	// defer db.Close()
	return nil

}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Person() *PersonRepository {
	if s.personRepository != nil {
		return s.personRepository
	}
	s.personRepository = &PersonRepository{store: s}

	return s.personRepository
}

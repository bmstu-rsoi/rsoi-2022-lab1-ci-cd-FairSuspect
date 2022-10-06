package store

import (
	"database/sql"

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
	str := "host=localhost dbname=persons user=program password=test port=5432 sslmode=disable"
	print("Conntecting to db with", str, "...\n")
	db, err := sql.Open("postgres", str)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

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

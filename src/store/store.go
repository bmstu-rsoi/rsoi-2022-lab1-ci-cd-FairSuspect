package store

import (
	"database/sql"
	"log"
	"os"

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
	dsn := os.Getenv("DATABASE_URL")
	if len(dsn) == 0 {
		dsn = "host=localhost dbname=persons user=program password=test port=5432 sslmode=disable"

	}
	log.Default().Println("Connecting to db... with " + dsn)

	db, err := sql.Open("postgres", dsn)

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

package store

import (
	"database/sql"
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
	dbURL := os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		dbURL = "localhost"
	}
	str := "host=" + dbURL + " dbname=das14gjflui68t user=winvrbuiddepcg password=740bc8cca39275cb85058dbdebb0777c2f223674ee57ed151e5fe62b1f0798b6 port=5432 sslmode=disable"
	print("Conntecting to db with ", str, "...\n")
	db, err := sql.Open("postgres", str)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	defer db.Close()
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

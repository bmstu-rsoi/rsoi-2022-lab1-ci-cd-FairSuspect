package store

import (
	"http-rest-api/internal/app/apiserver/model"
	"log"
)

type PersonRepository struct {
	store *Store
}

func (r *PersonRepository) Create(p *model.Person) (int, error) {
	id := -1
	err := r.store.db.QueryRow("INSERT INTO Persons (Name, Address, Work, Age) VALUES ($1, $2, $3, $4) RETURNING ID", p.Name, p.Address, p.Work, p.Age).Scan(&id)
	if err != nil {
		return -1, err
	}

	log.Default().Println(id)
	return id, nil
}
func (r *PersonRepository) GetAll() ([]*model.Person, error) {
	persons := []*model.Person{}
	query, err := r.store.db.Query("SELECT * from Persons")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		p := &model.Person{}
		if err := query.Scan(&p.ID, &p.Name, &p.Address, &p.Work, &p.Age); err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func (r *PersonRepository) GetById(id int) (*model.Person, error) {
	p := &model.Person{}
	if err := r.store.db.QueryRow("SELECT * from Persons WHERE Id = $1", id).Scan(&p.ID, &p.Name, &p.Address, &p.Work, &p.Age); err != nil {
		return nil, err
	}
	return p, nil
}
func (r *PersonRepository) DeleteById(id int) (*model.Person, error) {
	p := &model.Person{}
	if err := r.store.db.QueryRow("DELETE from Persons WHERE Id = $1", id).Scan(&p.ID, &p.Name, &p.Address, &p.Work, &p.Age); err != nil {
		return nil, err
	}
	return p, nil
}
func (r *PersonRepository) Patch(p *model.Person) (*model.Person, error) {
	if err := r.store.db.QueryRow("UPDATE Persons set Name = $2, Address = $3, Work = $4, Age = $5 WHERE Id = $1", p.ID, p.Name, p.Address, p.Work, p.Age).Scan(&p.ID); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PersonRepository) FindByName(name string) (*model.Person, error) {
	p := &model.Person{}
	if err := r.store.db.QueryRow("SELECT * from Persons WHERE name = $1", name).Scan(&p.ID, &p.Name, &p.Address, &p.Work, &p.Age); err != nil {
		return nil, err
	}
	return p, nil
}

package store

import (
	"fmt"
	"http-rest-api/internal/app/apiserver/model"
	"log"
	"strings"
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
	if _, err := r.store.db.Query("DELETE from Persons WHERE Id = $1", id); err != nil {
		return nil, err
	}
	return p, nil
}
func (r *PersonRepository) Patch(p *model.Person) (*model.Person, error) {
	var args string
	if len(p.Address) > 0 {
		args += "Address = '" + p.Address + "', "
	}
	if len(p.Name) > 0 {
		args += "Name = '" + p.Name + "', "
	}
	if len(p.Work) > 0 {
		args += "Work = '" + p.Work + "', "
	}
	if p.Age > 0 {
		args += fmt.Sprintf("Age = %d, ", p.Age)
	}
	args = strings.TrimRight(args, ", ")
	log.Default().Println(args)
	query := "UPDATE Persons set " + args + " WHERE Id = $1"
	log.Default().Println(query)

	if _, err := r.store.db.Query(query, p.ID); err != nil {
		return nil, err
	}
	p, err := r.GetById(p.ID)
	if err != nil {
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

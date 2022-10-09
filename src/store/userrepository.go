package store

import "http-rest-api/internal/app/apiserver/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow("INSER INTO users (name, password, datecreated) VALUES ($1, $2, NOW() RETURNING id", u.Name, u.Password).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByName(name string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT * from users WHERE name = $1", name).Scan(&u.Id, &u.Name, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

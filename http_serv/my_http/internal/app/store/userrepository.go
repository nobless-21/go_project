package store

import "my_http/internal/app/model"

type UserRepository struct{
	store *Store
}

func (u *UserRepository) Create(model.User) (error, model.User) {

}
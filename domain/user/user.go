package user

import "ecommerce/model"

type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Repository interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

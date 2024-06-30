package user

import (
	"ecommerce/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repository Repository
}

func New(r Repository) User {
	return User{repository: r}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()

	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUid", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword", err)
	}

	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.ID = ID
	m.Password = string(password)
	m.CreateAt = time.Now().Unix()

	err = u.repository.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	m.Password = ""
	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.repository.GetByEmail(email)

	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "repository.GetByEmail", err)
	}

	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.repository.GetAll()

	if err != nil {
		return model.Users{}, fmt.Errorf("%s %w", "reporsitory.GetAll", err)
	}

	return users, nil
}

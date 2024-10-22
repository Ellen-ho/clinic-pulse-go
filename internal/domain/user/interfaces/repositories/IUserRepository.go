package repositories

import (
	"clinic-pulse-go/internal/domain"
)

type UserRepository interface {
	FindById(id string) (*user.User, error)

	FindByEmail(email string) (*user.User, error)
	FindAdmin() (*user.User, error)
	Save(user *user.User) error
}


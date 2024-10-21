package repository

import (
	"my-go-backend/internal/domain"
	"strconv"
)

type UserRepository interface {
    GetByID(id int) (*domain.User, error)
    Create(user *domain.User) error
}

func NewUserRepository() UserRepository {
    return &userRepository{}
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
    return &domain.User{
        ID:        strconv.Itoa(id), 
        FirstName: "Tim",
        LastName:  "Lin",
        Role:      "DOCTOR", 
    }, nil
}

func (r *userRepository) Create(user *domain.User) error {
    return nil
}
package usecase

import (
    "my-go-backend/internal/domain"
    "my-go-backend/internal/repository"
)

type UserUsecase interface {
    GetUser(id int) (*domain.User, error)
    CreateUser(user *domain.User) error
}

type userUsecase struct {
    repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
    return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUser(id int) (*domain.User, error) {
    return u.repo.GetByID(id)
}

func (u *userUsecase) CreateUser(user *domain.User) error {
    return u.repo.Create(user)
}
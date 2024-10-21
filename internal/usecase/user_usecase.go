package usecase

import (
    "clinic-pulse-go/internal/domain"
    "clinic-pulse-go/internal/repository"
)

type UserUsecase interface {
    GetUser(id string) (*domain.User, error)  
    CreateUser(user *domain.User) error
}

type userUsecase struct {
    repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
    return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUser(id string) (*domain.User, error) {
    return u.repo.GetByID(id)  
}

func (u *userUsecase) CreateUser(user *domain.User) error {
    return u.repo.Create(user)
}
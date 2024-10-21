package repository

import (
    "gorm.io/gorm"
	"clinic-pulse-go/internal/domain"
)

type UserRepository interface {
    GetByID(id string) (*domain.User, error)  
    Create(user *domain.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}
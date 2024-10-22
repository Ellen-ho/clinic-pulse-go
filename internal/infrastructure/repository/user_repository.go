package repository

import (
	"time"
	"clinic-pulse-go/internal/shared"
	"clinic-pulse-go/internal/domain" 
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	Create(user *domain.User) error
}

type User struct {
	ID             string              `gorm:"type:uuid;primaryKey" json:"id"`
	Email          string              `gorm:"type:varchar(100);not null" json:"email"`
	HashedPassword string              `gorm:"type:varchar(255);not null" json:"-"`
	Role           shared.UserRoleType `gorm:"type:varchar(50);not null" json:"role"`
	CreatedAt      time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
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

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetHashedPassword() string {
	return u.HashedPassword
}

func (u *User) GetRole() shared.UserRoleType {
	return u.Role
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) UpdatePassword(newPassword string) {
	u.HashedPassword = newPassword
}

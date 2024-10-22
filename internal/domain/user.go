package domain

import (
	"time"
	"clinic-pulse-go/internal/shared"
)

type User struct {
    ID             string              `gorm:"type:uuid;primaryKey" json:"id"`
    Email          string              `gorm:"type:varchar(100);not null" json:"email"`
    HashedPassword string              `gorm:"type:varchar(255);not null" json:"-"`
    Role           shared.UserRoleType `gorm:"type:varchar(50);not null" json:"role"`
    CreatedAt      time.Time           `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt      time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
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

package domain

import (
	"time"
	"your_project/entity"
)

type IUserProps struct {
	ID             string
	Email          string
	HashedPassword string
	Role           entity.UserRoleType
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	Props IUserProps
}

func (u *User) GetID() string {
	return u.Props.ID
}

func (u *User) GetEmail() string {
	return u.Props.Email
}

func (u *User) GetHashedPassword() string {
	return u.Props.HashedPassword
}

func (u *User) GetRole() entity.UserRoleType {
	return u.Props.Role
}

func (u *User) GetCreatedAt() time.Time {
	return u.Props.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.Props.UpdatedAt
}

func (u *User) UpdatePassword(newPassword string) {
	u.Props.HashedPassword = newPassword
}

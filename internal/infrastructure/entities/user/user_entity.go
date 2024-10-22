package entity

import (
	"time"
	"gorm.io/gorm"
)

type UserRoleType string

const (
	Admin      UserRoleType = "ADMIN"
	Doctor     UserRoleType = "DOCTOR"
	Pharmacist UserRoleType = "PHARMACIST"
)

type UserEntity struct {
	ID        string       `gorm:"type:uuid;primaryKey"`
	Email     string       `gorm:"type:varchar(100);unique;not null"`
	Password  string       `gorm:"type:varchar(255);not null"`
	Role      UserRoleType `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
}

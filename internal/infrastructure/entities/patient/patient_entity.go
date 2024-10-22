package entity

import (
	"time"
	"gorm.io/gorm"
)

type GenderType string

const (
	Male      GenderType = "MALE"
	Female    GenderType = "FEMALE"
	NonBinary GenderType = "NON_BINARY"
)

type PatientEntity struct {
	ID        string     `gorm:"type:uuid;primaryKey"`
	FirstName string     `gorm:"type:varchar(50);not null"`
	LastName  string     `gorm:"type:varchar(50);not null"`
	FullName  string     `gorm:"type:varchar(50);default:'-';not null"`
	Gender    GenderType `gorm:"type:varchar(20);not null"`
	BirthDate time.Time  `gorm:"type:date;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

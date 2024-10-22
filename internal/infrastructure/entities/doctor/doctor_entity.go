package doctor

import (
	"time"
	"gorm.io/gorm"
	"clinic-pulse-go/internal/shared"
)

type DoctorEntity struct {
	ID             string     `gorm:"type:uuid;primaryKey"`
	Avatar         *string    `gorm:"type:varchar(255);nullable"`
	FirstName      string     `gorm:"type:varchar(50);not null"`
	LastName       string     `gorm:"type:varchar(50);not null"`
	Gender         GenderType `gorm:"type:varchar(20);not null"`
	BirthDate      time.Time  `gorm:"type:date;not null"`
	OnboardDate    time.Time  `gorm:"type:date;not null"`
	ResignationDate *time.Time `gorm:"type:date;nullable"`
	UserID         string     `gorm:"type:uuid;not null"`
	User           UserEntity `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}
package entity

import (
	"clinic-pulse-go/internal/shared"
	"gorm.io/gorm"
)

type ClinicEntity struct {
	ID      string          `gorm:"type:uuid;primaryKey"`
	Name    string          `gorm:"type:varchar(50);not null"`
	Address shared.IAddress `gorm:"type:jsonb;not null"`
}

func (ClinicEntity) TableName() string {
	return "clinics"
}
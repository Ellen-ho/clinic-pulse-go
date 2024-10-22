package entity

import (
	"time"
	"gorm.io/gorm"
)

type AcupunctureTreatmentEntity struct {
	ID            string     `gorm:"type:uuid;primaryKey"`
	StartAt       *time.Time `gorm:"type:timestamp;nullable"`
	EndAt         *time.Time `gorm:"type:timestamp;nullable"`
	BedID         *string    `gorm:"type:varchar(255);nullable"`
	AssignBedAt   *time.Time `gorm:"type:timestamp;nullable"`
	RemoveNeedleAt *time.Time `gorm:"type:timestamp;nullable"`
	NeedleCounts  *int       `gorm:"type:int;nullable"`
}

package timeSlot

import (
	"time"
	"gorm.io/gorm"
	"clinic-pulse-go/internal/shared"
)

type TimeSlotEntity struct {
	ID                 string                `gorm:"type:uuid;primaryKey"`
	StartAt            time.Time             `gorm:"type:timestamp;not null"`
	EndAt              time.Time             `gorm:"type:timestamp;not null"`
	TimePeriod         shared.TimePeriodType `gorm:"type:varchar(50);not null"`
	CreatedAt          time.Time             `gorm:"autoCreateTime"`
	UpdatedAt          time.Time             `gorm:"autoUpdateTime"`

	DoctorID           string                `gorm:"type:uuid;not null"`
	Doctor             DoctorEntity          `gorm:"foreignKey:DoctorID;references:ID"`

	ClinicID           string                `gorm:"type:uuid;not null"`
	Clinic             ClinicEntity          `gorm:"foreignKey:ClinicID;references:ID"`

	ConsultationRoomID string                `gorm:"type:uuid;not null"`
	ConsultationRoom   ConsultationRoomEntity `gorm:"foreignKey:ConsultationRoomID;references:ID"`
}
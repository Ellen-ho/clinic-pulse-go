package entity

import (
	"time"
	"gorm.io/gorm"
)

type TimePeriodType string

const (
	MorningSession   TimePeriodType = "MORNING_SESSION"
	AfternoonSession TimePeriodType = "AFTERNOON_SESSION"
	EveningSession   TimePeriodType = "EVENING_SESSION"
)

type TimeSlotEntity struct {
	ID                 string            `gorm:"type:uuid;primaryKey"`
	StartAt            time.Time         `gorm:"type:timestamp;not null"`
	EndAt              time.Time         `gorm:"type:timestamp;not null"`
	TimePeriod         TimePeriodType    `gorm:"type:varchar(50);not null"`
	CreatedAt          time.Time         `gorm:"autoCreateTime"`
	UpdatedAt          time.Time         `gorm:"autoUpdateTime"`

	DoctorID           string            `gorm:"type:uuid;not null"`
	Doctor             DoctorEntity      `gorm:"foreignKey:DoctorID;references:ID"`

	ClinicID           string            `gorm:"type:uuid;not null"`
	Clinic             ClinicEntity      `gorm:"foreignKey:ClinicID;references:ID"`

	ConsultationRoomID string            `gorm:"type:uuid;not null"`
	ConsultationRoom   ConsultationRoomEntity `gorm:"foreignKey:ConsultationRoomID;references:ID"`
}

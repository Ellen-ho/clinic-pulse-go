package entity

import (
	"time"
	"gorm.io/gorm"
)

type ConsultationSource string

const (
	OnlineBooking      ConsultationSource = "ONLINE_BOOKING"
	OnsiteRegistration ConsultationSource = "ONSITE_REGISTRATION"
)

type ConsultationStatus string

const (
	InConsultation                 ConsultationStatus = "IN_CONSULTATION"
	WaitingForConsultation          ConsultationStatus = "WAITING_FOR_CONSULTATION"
	WaitingForBedAssignment         ConsultationStatus = "WAITING_FOR_BED_ASSIGNMENT"
	WaitingForAcupunctureTreatment  ConsultationStatus = "WAITING_FOR_ACUPUNCTURE_TREATMENT"
	WaitingForNeedleRemoval         ConsultationStatus = "WAITING_FOR_NEEDLE_REMOVAL"
	WaitingForGetMedicine           ConsultationStatus = "WAITING_FOR_GET_MEDICINE"
	OnsiteCancel                    ConsultationStatus = "ONSITE_CANCEL"
	CheckOut                        ConsultationStatus = "CHECK_OUT"
	UndergoingAcupunctureTreatment  ConsultationStatus = "UNDERGOING_ACUPUNCTURE_TREATMENT"
)

type OnsiteCancelReasonType string

const (
	LongWaitingTime      OnsiteCancelReasonType = "LONG_WAITING_TIME"
	ServiceDissatisfaction OnsiteCancelReasonType = "SERVICE_DISSATISFACTION"
	PersonalEmergency      OnsiteCancelReasonType = "PERSONAL_EMERGENCY"
	NoParkingSpaces        OnsiteCancelReasonType = "NO_PARKING_SPACES"
)

type ConsultationEntity struct {
	ID                  string                 `gorm:"type:uuid;primaryKey"`
	Status              ConsultationStatus      `gorm:"type:varchar(255);default:WAITING_FOR_CONSULTATION"`
	Source              ConsultationSource      `gorm:"type:varchar(50);not null"`
	ConsultationNumber  int                    `gorm:"type:int"`
	CheckInAt           time.Time              `gorm:"type:timestamp"`
	StartAt             *time.Time             `gorm:"type:timestamp;nullable"`
	EndAt               *time.Time             `gorm:"type:timestamp;nullable"`
	CheckOutAt          *time.Time             `gorm:"type:timestamp;nullable"`
	OnsiteCancelAt      *time.Time             `gorm:"type:timestamp;nullable"`
	OnsiteCancelReason  *OnsiteCancelReasonType `gorm:"type:varchar(255);nullable"`
	IsFirstTimeVisit    bool                   `gorm:"type:boolean;default:false"`
	AcupunctureTreatmentID *string             `gorm:"type:uuid;nullable"`
	MedicineTreatmentID    *string             `gorm:"type:uuid;nullable"`
	PatientID           string                 `gorm:"type:uuid;not null"`
	TimeSlotID          string                 `gorm:"type:uuid;not null"`
	CreatedAt           time.Time              `gorm:"autoCreateTime"`
	UpdatedAt           time.Time              `gorm:"autoUpdateTime"`
}
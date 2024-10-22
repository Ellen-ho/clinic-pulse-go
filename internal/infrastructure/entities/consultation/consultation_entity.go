package consultation

import (
	"time"
	"clinic-pulse-go/internal/shared" 
)

type ConsultationEntity struct {
	ID                    string     `gorm:"type:uuid;primaryKey"`
	Status                shared.ConsultationStatus `gorm:"type:varchar(255);default:WAITING_FOR_CONSULTATION"` 
	Source                shared.ConsultationSource `gorm:"type:varchar(50);not null"`                         
	ConsultationNumber    int        `gorm:"type:int"`
	CheckInAt             time.Time  `gorm:"type:timestamp"`
	StartAt               *time.Time `gorm:"type:timestamp;nullable"`
	EndAt                 *time.Time `gorm:"type:timestamp;nullable"`
	CheckOutAt            *time.Time `gorm:"type:timestamp;nullable"`
	OnsiteCancelAt        *time.Time `gorm:"type:timestamp;nullable"`
	OnsiteCancelReason    *shared.OnsiteCancelReasonType `gorm:"type:varchar(255);nullable"`  
	IsFirstTimeVisit      bool       `gorm:"type:boolean;default:false"`
	PatientID             string     `gorm:"type:uuid;not null"`
	TimeSlotID            string     `gorm:"type:uuid;not null"`
	AcupunctureTreatment  *string    `gorm:"type:uuid;nullable"` 
	MedicineTreatment     *string    `gorm:"type:uuid;nullable"`  
	CreatedAt             time.Time  `gorm:"autoCreateTime"`
	UpdatedAt             time.Time  `gorm:"autoUpdateTime"`
}

func (ConsultationEntity) TableName() string {
    return "consultations" 
}
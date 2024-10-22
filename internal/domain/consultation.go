package domain

import (
	"time"
	"clinic-pulse-go/internal/shared"
)

type IConsultationProps struct {
	ID                   string
	Status               shared.ConsultationStatus 
	Source               shared.ConsultationSource  
	ConsultationNumber   int
	CheckInAt            time.Time
	StartAt              *time.Time
	EndAt                *time.Time
	CheckOutAt           *time.Time
	OnsiteCancelAt       *time.Time
	OnsiteCancelReason   *shared.OnsiteCancelReasonType  
	IsFirstTimeVisit     bool
	AcupunctureTreatment *AcupunctureTreatment
	MedicineTreatment    *MedicineTreatment
	PatientID            string
	TimeSlotID           string
}

type Consultation struct {
	Props IConsultationProps
}

func (c *Consultation) GetID() string {
	return c.Props.ID
}

func (c *Consultation) GetStatus() shared.ConsultationStatus {  
	return c.Props.Status
}

func (c *Consultation) GetSource() shared.ConsultationSource {  
	return c.Props.Source
}

func (c *Consultation) GetConsultationNumber() int {
	return c.Props.ConsultationNumber
}

func (c *Consultation) GetCheckInAt() time.Time {
	return c.Props.CheckInAt
}

func (c *Consultation) GetStartAt() *time.Time {
	return c.Props.StartAt
}

func (c *Consultation) GetEndAt() *time.Time {
	return c.Props.EndAt
}

func (c *Consultation) GetCheckOutAt() *time.Time {
	return c.Props.CheckOutAt
}

func (c *Consultation) GetOnsiteCancelAt() *time.Time {
	return c.Props.OnsiteCancelAt
}

func (c *Consultation) GetIsFirstTimeVisit() bool {
	return c.Props.IsFirstTimeVisit
}

func (c *Consultation) GetOnsiteCancelReason() *shared.OnsiteCancelReasonType {  
	return c.Props.OnsiteCancelReason
}

func (c *Consultation) GetPatientID() string {
	return c.Props.PatientID
}

func (c *Consultation) GetTimeSlotID() string {
	return c.Props.TimeSlotID
}

func (c *Consultation) GetAcupunctureTreatment() *AcupunctureTreatment {
	return c.Props.AcupunctureTreatment
}

func (c *Consultation) GetMedicineTreatment() *MedicineTreatment {
	return c.Props.MedicineTreatment
}

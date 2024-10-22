package domain

import (
	"time"
	"clinic-pulse-go/entity"
)

type IConsultationProps struct {
	ID                   string
	Status               entity.ConsultationStatus
	Source               entity.ConsultationSource
	ConsultationNumber   int
	CheckInAt            time.Time
	StartAt              *time.Time
	EndAt                *time.Time
	CheckOutAt           *time.Time
	OnsiteCancelAt       *time.Time
	OnsiteCancelReason   *entity.OnsiteCancelReasonType
	IsFirstTimeVisit     bool
	AcupunctureTreatment *entity.AcupunctureTreatmentEntity
	MedicineTreatment    *entity.MedicineTreatmentEntity
	PatientID            string
	TimeSlotID           string
}

type Consultation struct {
	Props IConsultationProps
}

func (c *Consultation) GetID() string {
	return c.Props.ID
}

func (c *Consultation) GetStatus() entity.ConsultationStatus {
	return c.Props.Status
}

func (c *Consultation) GetSource() entity.ConsultationSource {
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

func (c *Consultation) GetOnsiteCancelReason() *entity.OnsiteCancelReasonType {
	return c.Props.OnsiteCancelReason
}

func (c *Consultation) GetAcupunctureTreatment() *entity.AcupunctureTreatmentEntity {
	return c.Props.AcupunctureTreatment
}

func (c *Consultation) GetMedicineTreatment() *entity.MedicineTreatmentEntity {
	return c.Props.MedicineTreatment
}

func (c *Consultation) GetPatientID() string {
	return c.Props.PatientID
}

func (c *Consultation) GetTimeSlotID() string {
	return c.Props.TimeSlotID
}

func (c *Consultation) UpdateStartAt(status entity.ConsultationStatus, startAt time.Time) {
	c.Props.Status = status
	c.Props.StartAt = &startAt
}

func (c *Consultation) UpdateToWaitAcupuncture(status entity.ConsultationStatus) {
	c.Props.Status = status
}

func (c *Consultation) UpdateToWaitForBed(status entity.ConsultationStatus, endAt time.Time, acupunctureTreatment *entity.AcupunctureTreatmentEntity) {
	c.Props.Status = status
	c.Props.EndAt = &endAt
	c.Props.AcupunctureTreatment = acupunctureTreatment
}

func (c *Consultation) UpdateToMedicine(status entity.ConsultationStatus, endAt time.Time, medicineTreatment *entity.MedicineTreatmentEntity) {
	c.Props.Status = status
	c.Props.EndAt = &endAt
	c.Props.MedicineTreatment = medicineTreatment
}

func (c *Consultation) UpdateToCheckOutAt(status entity.ConsultationStatus, checkOutAt time.Time) {
	c.Props.Status = status
	c.Props.CheckOutAt = &checkOutAt
}

func (c *Consultation) UpdateToOnsiteCancel(status entity.ConsultationStatus, onsiteCancelAt time.Time, onsiteCancelReason entity.OnsiteCancelReasonType) {
	c.Props.Status = status
	c.Props.OnsiteCancelAt = &onsiteCancelAt
	c.Props.OnsiteCancelReason = &onsiteCancelReason
}

func (c *Consultation) UpdateStatus(status entity.ConsultationStatus) {
	c.Props.Status = status
}
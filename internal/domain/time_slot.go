package domain

import (
	"time"
	"clinic-pulse-go/entity"
)

type ITimeSlotProps struct {
	ID                 string
	StartAt            time.Time
	EndAt              time.Time
	TimePeriod         entity.TimePeriodType
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DoctorID           string
	ClinicID           string
	ConsultationRoomID string
}

type TimeSlot struct {
	Props ITimeSlotProps
}

func (t *TimeSlot) GetID() string {
	return t.Props.ID
}

func (t *TimeSlot) GetStartAt() time.Time {
	return t.Props.StartAt
}

func (t *TimeSlot) GetEndAt() time.Time {
	return t.Props.EndAt
}

func (t *TimeSlot) GetTimePeriod() entity.TimePeriodType {
	return t.Props.TimePeriod
}

func (t *TimeSlot) GetCreatedAt() time.Time {
	return t.Props.CreatedAt
}

func (t *TimeSlot) GetUpdatedAt() time.Time {
	return t.Props.UpdatedAt
}

func (t *TimeSlot) GetDoctorID() string {
	return t.Props.DoctorID
}

func (t *TimeSlot) GetClinicID() string {
	return t.Props.ClinicID
}

func (t *TimeSlot) GetConsultationRoomID() string {
	return t.Props.ConsultationRoomID
}
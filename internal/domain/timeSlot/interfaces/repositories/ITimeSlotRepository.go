package repositories

import (
	"time"
	"clinic-pulse-go/internal/domain"
)

type TimeSlotRepository interface {

	GetById(id string) (*timeslot.TimeSlot, error)

	FindByClinicId(clinicId string) ([]ClinicTimeSlotResult, error)

	FindByDoctorId(doctorId string) ([]DoctorTimeSlotResult, error)
}

type TimeSlotDurationResult struct {
	TotalTimeSlots int
	Data           []TimeSlotCountByDate
}

type TimeSlotCountByDate struct {
	Date          string
	TimeSlotCount int
}

type TimeSlotIDResult struct {
	TimeSlotID string
}

type CurrentTimeSlotResult struct {
	TimeSlotID           string
	ClinicID             string
	ConsultationRoomNumber consultationroom.RoomNumberType
	TimePeriod           timeslot.TimePeriodType
}

type AdminTimeSlotResult struct {
	ID                    string
	ClinicID              string
	ConsultationRoomNumber consultationroom.RoomNumberType
	TimePeriod            timeslot.TimePeriodType
}

type ClinicTimeSlotResult struct {
	ID            string
	StartAt       time.Time
	EndAt         time.Time
	TimePeriod    timeslot.TimePeriodType
	ClinicID      string
	Doctor        DoctorInfo
	ConsultationRoom ConsultationRoomInfo
}

type DoctorTimeSlotResult struct {
	ID            string
	StartAt       time.Time
	EndAt         time.Time
	TimePeriod    timeslot.TimePeriodType
	ClinicID      string
	Doctor        DoctorInfo
	ConsultationRoom ConsultationRoomInfo
}

type DoctorInfo struct {
	ID        string
	FirstName string
	LastName  string
}

type ConsultationRoomInfo struct {
	ID         string
	RoomNumber consultationroom.RoomNumberType
}


package timeSlot

import (
	"clinic-pulse-go/internal/domain"
)

type TimeSlotMapper struct{}

func (m *TimeSlotMapper) ToDomainModel(e *timeSlot.TimeSlotEntity) *domain.TimeSlot {
	return &domain.TimeSlot{
		Props: domain.ITimeSlotProps{
			ID:                 e.ID,
			StartAt:            e.StartAt,
			EndAt:              e.EndAt,
			TimePeriod:         e.TimePeriod,
			CreatedAt:          e.CreatedAt,
			UpdatedAt:          e.UpdatedAt,
			DoctorID:           e.DoctorID,
			ClinicID:           e.ClinicID,
			ConsultationRoomID: e.ConsultationRoomID,
		},
	}
}

func (m *TimeSlotMapper) ToEntity(d *domain.TimeSlot) *timeSlot.TimeSlotEntity {
	return &timeSlot.TimeSlotEntity{
		ID:                 d.GetID(),
		StartAt:            d.GetStartAt(),
		EndAt:              d.GetEndAt(),
		TimePeriod:         d.GetTimePeriod(),
		DoctorID:           d.GetDoctorID(),
		ClinicID:           d.GetClinicID(),
		ConsultationRoomID: d.GetConsultationRoomID(),
	}
}

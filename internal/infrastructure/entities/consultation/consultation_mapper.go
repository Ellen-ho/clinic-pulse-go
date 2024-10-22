package mapper

import (
	"clinic-pulse-go/domain"
	"clinic-pulse-go/entity"
)

type ConsultationMapper struct{}

func (m *ConsultationMapper) ToDomainModel(e *entity.ConsultationEntity) *domain.Consultation {
	return &domain.Consultation{
		Props: domain.IConsultationProps{
			ID:                 e.ID,
			Status:             e.Status,
			Source:             e.Source,
			ConsultationNumber: e.ConsultationNumber,
			CheckInAt:          e.CheckInAt,
			StartAt:            e.StartAt,
			EndAt:              e.EndAt,
			CheckOutAt:         e.CheckOutAt,
			OnsiteCancelAt:     e.OnsiteCancelAt,
			OnsiteCancelReason: e.OnsiteCancelReason,
			IsFirstTimeVisit:   e.IsFirstTimeVisit,
			PatientID:          e.PatientID,
			TimeSlotID:         e.TimeSlotID,
			AcupunctureTreatment: e.AcupunctureTreatment,
			MedicineTreatment:    e.MedicineTreatment,
		},
	}
}

func (m *ConsultationMapper) ToEntity(d *domain.Consultation) *entity.ConsultationEntity {
	return &entity.ConsultationEntity{
		ID:                 d.GetID(),
		Status:             d.GetStatus(),
		Source:             d.GetSource(),
		ConsultationNumber: d.GetConsultationNumber(),
		CheckInAt:          d.GetCheckInAt(),
		StartAt:            d.GetStartAt(),
		EndAt:              d.GetEndAt(),
		CheckOutAt:         d.GetCheckOutAt(),
		OnsiteCancelAt:     d.GetOnsiteCancelAt(),
		OnsiteCancelReason: d.GetOnsiteCancelReason(),
		IsFirstTimeVisit:   d.GetIsFirstTimeVisit(),
		PatientID:          d.GetPatientID(),
		TimeSlotID:         d.GetTimeSlotID(),
		AcupunctureTreatment: d.GetAcupunctureTreatment(),
		MedicineTreatment:    d.GetMedicineTreatment(),
	}
}

	

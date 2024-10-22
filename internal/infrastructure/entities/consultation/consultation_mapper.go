package consultation

import (
	"clinic-pulse-go/internal/domain"
	"clinic-pulse-go/internal/shared"
)

type ConsultationMapper struct{}

func (m *ConsultationMapper) ToDomainModel(e *ConsultationEntity) *domain.Consultation {
	return &domain.Consultation{
		Props: domain.IConsultationProps{
			ID:                 e.ID,
			Status:             shared.ConsultationStatus(e.Status),
			Source:             shared.ConsultationSource(e.Source),
			ConsultationNumber: e.ConsultationNumber,
			CheckInAt:          e.CheckInAt,
			StartAt:            e.StartAt,
			EndAt:              e.EndAt,
			CheckOutAt:         e.CheckOutAt,
			OnsiteCancelAt:     e.OnsiteCancelAt,
			OnsiteCancelReason: (*shared.OnsiteCancelReasonType)(e.OnsiteCancelReason),
			IsFirstTimeVisit:   e.IsFirstTimeVisit,
			PatientID:          e.PatientID,
			TimeSlotID:         e.TimeSlotID,
			AcupunctureTreatment: &domain.AcupunctureTreatment{
				Props: domain.IAcupunctureTreatmentProps{ID: *e.AcupunctureTreatment},
			}, 
			MedicineTreatment: &domain.MedicineTreatment{
				Props: domain.IMedicineTreatmentProps{ID: *e.MedicineTreatment},
			},
		},
	}
}

func (m *ConsultationMapper) ToEntity(d *domain.Consultation) *ConsultationEntity {
	return &ConsultationEntity{
		ID:                 d.GetID(),
		Status:             shared.ConsultationStatus(d.GetStatus()),
		Source:             shared.ConsultationSource(d.GetSource()),
		ConsultationNumber: d.GetConsultationNumber(),
		CheckInAt:          d.GetCheckInAt(),
		StartAt:            d.GetStartAt(),
		EndAt:              d.GetEndAt(),
		CheckOutAt:         d.GetCheckOutAt(),
		OnsiteCancelAt:     d.GetOnsiteCancelAt(),
		OnsiteCancelReason: (*shared.OnsiteCancelReasonType)(d.GetOnsiteCancelReason()),
		IsFirstTimeVisit:   d.GetIsFirstTimeVisit(),
		PatientID:          d.GetPatientID(),
		TimeSlotID:         d.GetTimeSlotID(),
		AcupunctureTreatment: &(*d.GetAcupunctureTreatment()).Props.ID, 
		MedicineTreatment:    &(*d.GetMedicineTreatment()).Props.ID,   
	}
}

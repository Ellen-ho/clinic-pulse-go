package treatment

import (
	"clinic-pulse-go/internal/domain"
)

type AcupunctureTreatmentMapper struct{}

func (m *AcupunctureTreatmentMapper) ToDomainModel(e *treatment.AcupunctureTreatmentEntity) *domain.AcupunctureTreatment {
	return &domain.AcupunctureTreatment{
		Props: domain.IAcupunctureTreatmentProps{
			ID:            e.ID,
			StartAt:       e.StartAt,
			EndAt:         e.EndAt,
			BedID:         e.BedID,
			AssignBedAt:   e.AssignBedAt,
			RemoveNeedleAt: e.RemoveNeedleAt,
			NeedleCounts:  e.NeedleCounts,
		},
	}
}

func (m *AcupunctureTreatmentMapper) ToEntity(d *domain.AcupunctureTreatment) *treatment.AcupunctureTreatmentEntity {
	return &treatment.AcupunctureTreatmentEntity{
		ID:            d.GetID(),
		StartAt:       d.GetStartAt(),
		EndAt:         d.GetEndAt(),
		BedID:         d.GetBedID(),
		AssignBedAt:   d.GetAssignBedAt(),
		RemoveNeedleAt: d.GetRemoveNeedleAt(),
		NeedleCounts:  d.GetNeedleCounts(),
	}
}

package mapper

import (
	"clinic-pulse-go/domain"
	"clinic-pulse-go/entity"
)

type AcupunctureTreatmentMapper struct{}

func (m *AcupunctureTreatmentMapper) ToDomainModel(e *entity.AcupunctureTreatmentEntity) *domain.AcupunctureTreatment {
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

func (m *AcupunctureTreatmentMapper) ToEntity(d *domain.AcupunctureTreatment) *entity.AcupunctureTreatmentEntity {
	return &entity.AcupunctureTreatmentEntity{
		ID:            d.GetID(),
		StartAt:       d.GetStartAt(),
		EndAt:         d.GetEndAt(),
		BedID:         d.GetBedID(),
		AssignBedAt:   d.GetAssignBedAt(),
		RemoveNeedleAt: d.GetRemoveNeedleAt(),
		NeedleCounts:  d.GetNeedleCounts(),
	}
}

package treatment

import (
	"clinic-pulse-go/internal/domain"
)

type MedicineTreatmentMapper struct{}

func (m *MedicineTreatmentMapper) ToDomainModel(e *treatment.MedicineTreatmentEntity) *domain.MedicineTreatment {
	return &domain.MedicineTreatment{
		Props: domain.IMedicineTreatmentProps{
			ID:            e.ID,
			GetMedicineAt: e.GetMedicineAt,
		},
	}
}

func (m *MedicineTreatmentMapper) ToEntity(d *domain.MedicineTreatment) *treatment.MedicineTreatmentEntity {
	return &treatment.MedicineTreatmentEntity{
		ID:            d.GetID(),
		GetMedicineAt: d.GetGetMedicineAt(),
	}
}

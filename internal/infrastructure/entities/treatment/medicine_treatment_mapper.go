package mapper

import (
	"clinic-pulse-go/domain"
	"clinic-pulse-go/entity"
)

type MedicineTreatmentMapper struct{}

func (m *MedicineTreatmentMapper) ToDomainModel(e *entity.MedicineTreatmentEntity) *domain.MedicineTreatment {
	return &domain.MedicineTreatment{
		Props: domain.IMedicineTreatmentProps{
			ID:            e.ID,
			GetMedicineAt: e.GetMedicineAt,
		},
	}
}

func (m *MedicineTreatmentMapper) ToEntity(d *domain.MedicineTreatment) *entity.MedicineTreatmentEntity {
	return &entity.MedicineTreatmentEntity{
		ID:            d.GetID(),
		GetMedicineAt: d.GetGetMedicineAt(),
	}
}

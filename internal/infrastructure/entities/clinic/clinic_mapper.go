package mapper

import (
	"clinic-pulse-go/internal/domain/clinic"
	"clinic-pulse-go/internal/shared"
)

type ClinicMapper struct{}

func (m *ClinicMapper) ToDomainModel(entity entities.ClinicEntity) clinic.Clinic {
	return clinic.Clinic{
		Props: clinic.IClinicProps{
			ID:      entity.ID,
			Name:    entity.Name,
			Address: entity.Address, 
		},
	}
}

func (m *ClinicMapper) ToPersistence(domainModel clinic.Clinic) entities.ClinicEntity {
	return entities.ClinicEntity{
		ID:      domainModel.GetID(),
		Name:    domainModel.GetName(),
		Address: domainModel.GetAddress(), 
	}
}

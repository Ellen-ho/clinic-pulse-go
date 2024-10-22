package patient

import (
	"clinic-pulse-go/internal/domain"
)

type PatientMapper struct{}

func (m *PatientMapper) ToDomainModel(e *patient.PatientEntity) *domain.Patient {
	return &domain.Patient{
		Props: domain.IPatientProps{
			ID:        e.ID,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			FullName:  e.FullName,
			BirthDate: e.BirthDate,
			Gender:    e.Gender,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
	}
}

func (m *PatientMapper) ToEntity(d *domain.Patient) *patient.PatientEntity {
	return &patient.PatientEntity{
		ID:        d.GetID(),
		FirstName: d.GetFirstName(),
		LastName:  d.GetLastName(),
		FullName:  d.GetFullName(),
		BirthDate: d.GetBirthDate(),
		Gender:    d.GetGender(),
		CreatedAt: d.GetCreatedAt(),
		UpdatedAt: d.GetUpdatedAt(),
	}
}
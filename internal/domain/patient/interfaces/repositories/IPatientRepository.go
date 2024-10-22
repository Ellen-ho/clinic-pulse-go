package repositories

import (
	"clinic-pulse-go/internal/domain"
)

type PatientRepository interface {
	FindByName(searchText string) ([]PatientBasicInfo, error)

	Save(patient *patient.Patient) error
}

type PatientBasicInfo struct {
	ID       string
	FullName string
}


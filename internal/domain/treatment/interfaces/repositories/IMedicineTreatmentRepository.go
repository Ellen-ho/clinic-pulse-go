package repositories

import (
	"clinic-pulse-go/internal/domain"
)

type MedicineTreatmentRepository interface {
	Save(medicineTreatment *treatment.MedicineTreatment) error

	GetById(id string) (*treatment.MedicineTreatment, error)

	FindByConsultationId(consultationId string) (*treatment.MedicineTreatment, error)
}


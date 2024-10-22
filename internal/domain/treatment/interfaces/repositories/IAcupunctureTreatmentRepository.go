package repositories

import (
	"clinic-pulse-go/internal/domain"
)

type AcupunctureTreatmentRepository interface {
	Save(acupunctureTreatment *treatment.AcupunctureTreatment) error

	GetById(id string) (*treatment.AcupunctureTreatment, error)

	FindByConsultationId(consultationId string) (*treatment.AcupunctureTreatment, error)
}

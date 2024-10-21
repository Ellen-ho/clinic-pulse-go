package repository

import (
	"my-go-backend/internal/domain"
)

type IConsultationRepository interface {
	FindByQuery(limit, offset int, startDate, endDate, clinicId, timePeriod string, totalDurationMin, totalDurationMax int, patientName, patientId, doctorId string) ([]domain.Consultation, int, error)
}

func NewConsultationRepository() *ConsultationRepository {
	return &ConsultationRepository{}
}

func (r *ConsultationRepository) FindByQuery(limit, offset int, startDate, endDate, clinicId, timePeriod string, totalDurationMin, totalDurationMax int, patientName, patientId, doctorId string) ([]domain.Consultation, int, error) {
	consultations := []domain.Consultation{
		{
			ID:                "1",
			ConsultationNumber: 123,
			ConsultationDate:   "2024-10-01",
			TotalDuration:      30,
			Doctor:             domain.Doctor{FirstName: "Tim", LastName: "Lin"},
			Patient:            domain.Patient{FirstName: "Jane", LastName: "Su", Gender: "Female", Age: 30},
		},
	}

	return consultations, 1, nil 
}
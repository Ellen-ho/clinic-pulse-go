package repository

import "my-go-backend/internal/domain"

type DoctorRepository struct {
}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{}
}

func (r *DoctorRepository) FindByUserId(userID string) (*domain.Doctor, error) {
	return &domain.Doctor{ID: "doctor-123", FirstName: "Tim", LastName: "Lin"}, nil
}
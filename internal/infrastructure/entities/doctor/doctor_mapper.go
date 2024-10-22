package mapper

import (
	"clinic-pulse-go/domain"
	"clinic-pulse-go/entity"
)

type DoctorMapper struct{}

func (m *DoctorMapper) ToDomainModel(e *entity.DoctorEntity) *domain.Doctor {
	return &domain.Doctor{
		Props: domain.IDoctorProps{
			ID:             e.ID,
			Avatar:         e.Avatar,
			FirstName:      e.FirstName,
			LastName:       e.LastName,
			Gender:         e.Gender,
			BirthDate:      e.BirthDate,
			OnboardDate:    e.OnboardDate,
			ResignationDate: e.ResignationDate,
			User:           &e.User,
		},
	}
}

func (m *DoctorMapper) ToEntity(d *domain.Doctor) *entity.DoctorEntity {
	return &entity.DoctorEntity{
		ID:             d.GetID(),
		Avatar:         d.GetAvatar(),
		FirstName:      d.GetFirstName(),
		LastName:       d.GetLastName(),
		Gender:         d.GetGender(),
		BirthDate:      d.GetBirthDate(),
		OnboardDate:    d.GetOnboardDate(),
		ResignationDate: d.GetResignationDate(),
		UserID:         d.GetUser().ID,
	}
}
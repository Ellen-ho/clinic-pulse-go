package mapper

import (
	"clinic-pulse-go/domain"
	"clinic-pulse-go/entity"
)

type UserMapper struct{}

func (m *UserMapper) ToDomainModel(e *entity.UserEntity) *domain.User {
	return &domain.User{
		Props: domain.IUserProps{
			ID:             e.ID,
			Email:          e.Email,
			HashedPassword: e.Password,
			Role:           e.Role,
			CreatedAt:      e.CreatedAt,
			UpdatedAt:      e.UpdatedAt,
		},
	}
}

func (m *UserMapper) ToEntity(d *domain.User) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        d.GetID(),
		Email:     d.GetEmail(),
		Password:  d.GetHashedPassword(),
		Role:      d.GetRole(),
		CreatedAt: d.GetCreatedAt(),
		UpdatedAt: d.GetUpdatedAt(),
	}
}

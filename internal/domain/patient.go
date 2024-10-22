package domain

import (
	"time"
)

type GenderType string

const (
	Male      GenderType = "MALE"
	Female    GenderType = "FEMALE"
	NonBinary GenderType = "NON_BINARY"
)

type IPatientProps struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	BirthDate time.Time
	Gender    GenderType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Patient struct {
	Props IPatientProps
}

func (p *Patient) GetID() string {
	return p.Props.ID
}

func (p *Patient) GetFirstName() string {
	return p.Props.FirstName
}

func (p *Patient) GetLastName() string {
	return p.Props.LastName
}

func (p *Patient) GetFullName() string {
	return p.Props.FullName
}

func (p *Patient) GetBirthDate() time.Time {
	return p.Props.BirthDate
}

func (p *Patient) GetGender() GenderType {
	return p.Props.Gender
}

func (p *Patient) GetCreatedAt() time.Time {
	return p.Props.CreatedAt
}

func (p *Patient) GetUpdatedAt() time.Time {
	return p.Props.UpdatedAt
}

type IPatientProfileUpdateData struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	Gender    GenderType
}

func (p *Patient) UpdateData(data IPatientProfileUpdateData) {
	p.Props.FirstName = data.FirstName
	p.Props.LastName = data.LastName
	p.Props.BirthDate = data.BirthDate
	p.Props.Gender = data.Gender
}
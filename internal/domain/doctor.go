package domain

import (
	"time"
	"clinic-pulse-go/internal/shared"
)

type IDoctorProps struct {
	ID              string
	Avatar          *string
	FirstName       string
	LastName        string
	Gender          shared.GenderType  
	BirthDate       time.Time
	OnboardDate     time.Time
	ResignationDate *time.Time
	User            *shared.User      
}

type Doctor struct {
	Props IDoctorProps
}

func (d *Doctor) GetID() string {
	return d.Props.ID
}

func (d *Doctor) GetAvatar() *string {
	return d.Props.Avatar
}

func (d *Doctor) GetFirstName() string {
	return d.Props.FirstName
}

func (d *Doctor) GetLastName() string {
	return d.Props.LastName
}

func (d *Doctor) GetGender() shared.GenderType {  
	return d.Props.Gender
}

func (d *Doctor) GetBirthDate() time.Time {
	return d.Props.BirthDate
}

func (d *Doctor) GetOnboardDate() time.Time {
	return d.Props.OnboardDate
}

func (d *Doctor) GetResignationDate() *time.Time {
	return d.Props.ResignationDate
}

func (d *Doctor) GetUser() *shared.User {  
	return d.Props.User
}

func (d *Doctor) UpdateAvatar(newAvatar string) {
	d.Props.Avatar = &newAvatar
}

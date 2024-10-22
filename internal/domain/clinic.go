package domain

import (
	"clinic-pulse-go/internal/shared"
)

type IClinicProps struct {
	ID      string
	Name    string
	Address shared.IAddress `json:"address"` 
}

type Clinic struct {
	props IClinicProps
}

func NewClinic(props IClinicProps) *Clinic {
	return &Clinic{props: props}
}

func (c *Clinic) ID() string {
	return c.props.ID
}

func (c *Clinic) Name() string {
	return c.props.Name
}

func (c *Clinic) Address() shared.IAddress {
	return c.props.Address
}

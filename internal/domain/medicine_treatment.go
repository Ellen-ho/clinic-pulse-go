package domain

import (
	"time"
)

type IMedicineTreatmentProps struct {
	ID            string
	GetMedicineAt *time.Time
}

type MedicineTreatment struct {
	Props IMedicineTreatmentProps
}

func (m *MedicineTreatment) GetID() string {
	return m.Props.ID
}

func (m *MedicineTreatment) GetGetMedicineAt() *time.Time {
	return m.Props.GetMedicineAt
}

type IMedicineTreatmentUpdate struct {
	GetMedicineAt time.Time
}

func (m *MedicineTreatment) UpdateMedicineTreatment(data IMedicineTreatmentUpdate) {
	m.Props.GetMedicineAt = &data.GetMedicineAt
}

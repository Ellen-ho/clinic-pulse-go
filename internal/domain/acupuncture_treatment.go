package domain

import (
	"time"
)

type IAcupunctureTreatmentProps struct {
	ID            string
	StartAt       *time.Time
	EndAt         *time.Time
	BedID         *string
	AssignBedAt   *time.Time
	RemoveNeedleAt *time.Time
	NeedleCounts  *int
}

type AcupunctureTreatment struct {
	Props IAcupunctureTreatmentProps
}

func (a *AcupunctureTreatment) GetID() string {
	return a.Props.ID
}

func (a *AcupunctureTreatment) GetStartAt() *time.Time {
	return a.Props.StartAt
}

func (a *AcupunctureTreatment) GetEndAt() *time.Time {
	return a.Props.EndAt
}

func (a *AcupunctureTreatment) GetBedID() *string {
	return a.Props.BedID
}

func (a *AcupunctureTreatment) GetAssignBedAt() *time.Time {
	return a.Props.AssignBedAt
}

func (a *AcupunctureTreatment) GetRemoveNeedleAt() *time.Time {
	return a.Props.RemoveNeedleAt
}

func (a *AcupunctureTreatment) GetNeedleCounts() *int {
	return a.Props.NeedleCounts
}

type IAcupunctureTreatmentAssignBedUpdate struct {
	AssignBedAt time.Time
	BedID       string
}

func (a *AcupunctureTreatment) UpdateAssignBed(data IAcupunctureTreatmentAssignBedUpdate) {
	a.Props.AssignBedAt = &data.AssignBedAt
	a.Props.BedID = &data.BedID
}

type IAcupunctureTreatmentStartAtUpdate struct {
	StartAt      time.Time
	EndAt        time.Time
	NeedleCounts int
}

func (a *AcupunctureTreatment) UpdateStartAt(data IAcupunctureTreatmentStartAtUpdate) {
	a.Props.StartAt = &data.StartAt
	a.Props.EndAt = &data.EndAt
	a.Props.NeedleCounts = &data.NeedleCounts
}

type IAcupunctureTreatmentRemoveNeedleAtUpdate struct {
	RemoveNeedleAt time.Time
}

func (a *AcupunctureTreatment) UpdateRemoveNeedleAt(data IAcupunctureTreatmentRemoveNeedleAtUpdate) {
	a.Props.RemoveNeedleAt = &data.RemoveNeedleAt
}

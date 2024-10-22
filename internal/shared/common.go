package shared

import (
    "time"
)

type ConsultationQueryResult struct {
	TotalCount    int                  
	Consultations []ConsultationDetail 
}

type ConsultationDetail struct {
    ID                   string
    Status               ConsultationStatus
    Source               ConsultationSource
    ConsultationNumber   int
    CheckInAt            time.Time
    StartAt              *time.Time
    EndAt                *time.Time
    CheckOutAt           *time.Time
    OnsiteCancelAt       *time.Time
    OnsiteCancelReason   *OnsiteCancelReasonType
    IsFirstTimeVisit     bool
    PatientID            string
    TimeSlotID           string
    AcupunctureTreatment *AcupunctureTreatment
    MedicineTreatment    *MedicineTreatment
}


type ConsultationStatus string

const (
	InConsultation             ConsultationStatus = "IN_CONSULTATION"
	WaitingForConsultation     ConsultationStatus = "WAITING_FOR_CONSULTATION"
	WaitingForBedAssignment    ConsultationStatus = "WAITING_FOR_BED_ASSIGNMENT"
	WaitingForAcupuncture      ConsultationStatus = "WAITING_FOR_ACUPUNCTURE_TREATMENT"
	WaitingForNeedleRemoval    ConsultationStatus = "WAITING_FOR_NEEDLE_REMOVAL"
	WaitingForMedicine         ConsultationStatus = "WAITING_FOR_MEDICINE"
	OnsiteCancel               ConsultationStatus = "ONSITE_CANCEL"
	CheckOut                   ConsultationStatus = "CHECK_OUT"
	UndergoingAcupuncture      ConsultationStatus = "UNDERGOING_ACUPUNCTURE_TREATMENT"
)

type ConsultationSource string

const (
	OnlineBooking        ConsultationSource = "ONLINE_BOOKING"
	OnsiteRegistration   ConsultationSource = "ONSITE_REGISTRATION"
)

type OnsiteCancelReasonType string

const (
	LongWaitingTime      OnsiteCancelReasonType = "LONG_WAITING_TIME"
	ServiceDissatisfaction OnsiteCancelReasonType = "SERVICE_DISSATISFACTION"
	PersonalEmergency    OnsiteCancelReasonType = "PERSONAL_EMERGENCY"
	NoParkingSpaces      OnsiteCancelReasonType = "NO_PARKING_SPACES"
)

type GenderType string

const (
	Male      GenderType = "MALE"
	Female    GenderType = "FEMALE"
	NonBinary GenderType = "NON_BINARY"
)

type TimePeriodType string

const (
	MorningSession   TimePeriodType = "MORNING_SESSION"
	AfternoonSession TimePeriodType = "AFTERNOON_SESSION"
	EveningSession   TimePeriodType = "EVENING_SESSION"
)

type UserRoleType string

const (
    Admin      UserRoleType = "ADMIN"
    Doctor     UserRoleType = "DOCTOR"
)

type User struct {
	ID             string
	Email          string
	HashedPassword string
	Role           UserRoleType
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type IAddress struct {
	Line1        string  `json:"line1"`
	Line2        *string `json:"line2,omitempty"`
	City         string  `json:"city"`
	StateProvince *string `json:"stateProvince,omitempty"`
	PostalCode   *string `json:"postalCode,omitempty"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"countryCode"`
}

type AcupunctureTreatment struct {
	ID            string
	StartAt       *time.Time
	EndAt         *time.Time
	NeedleCounts  *int
	AssignBedAt   *time.Time
	BedID         *string
	RemoveNeedleAt *time.Time
}

type MedicineTreatment struct {
	ID            string
	GetMedicineAt *time.Time
}

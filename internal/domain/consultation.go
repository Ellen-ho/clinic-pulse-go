package domain

type GenderType string

type TimePeriodType string

type OnsiteCancelReasonType string

type TreatmentType string

type Doctor struct {
	ID        string
	FirstName string
	LastName  string
}

type Patient struct {
	FirstName string
	LastName  string
	Gender    GenderType
	Age       int
}

type Consultation struct {
	ID                   string
	IsOnsiteCanceled      bool
	OnsiteCancelReason    OnsiteCancelReasonType
	ConsultationNumber    int
	ConsultationDate      string
	ConsultationTimePeriod TimePeriodType
	Doctor                Doctor
	Patient               Patient
	TreatmentType         TreatmentType
	TotalDuration         int
}

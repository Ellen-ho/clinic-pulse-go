package repositories

import (
	"clinic-pulse-go/internal/shared"
)

type IConsultationRepository interface {
	FindByQuery(
        limit int,
        offset int,
        startDate string,
        endDate string,
        clinicId *string,
        timePeriod *string,
        totalDurationMin *int,
        totalDurationMax *int,
        patientName *string,
        patientId *string,
        doctorId *string,
    ) (*shared.ConsultationQueryResult, error)
}
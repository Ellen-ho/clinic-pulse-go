package repository

import (
    "time"
    "clinic-pulse-go/internal/shared"
    "clinic-pulse-go/internal/infrastructure/entities/consultation"
    "gorm.io/gorm"
)

type ConsultationRepository struct {
    db *gorm.DB
}

func NewConsultationRepository(db *gorm.DB) *ConsultationRepository {
    return &ConsultationRepository{db: db}
}

func (r *ConsultationRepository) FindByQuery(
    limit, offset int,
    startDate, endDate string,
    clinicId, timePeriod *string,
    totalDurationMin, totalDurationMax *int,
    patientName, patientId, doctorId *string,
) (*shared.ConsultationQueryResult, error) {
    var consultations []consultation.ConsultationEntity
    var totalCount int64

    startDateTime, err := time.Parse("2006-01-02", startDate)
    if err != nil {
        return nil, err
    }
    endDateTime, err := time.Parse("2006-01-02", endDate)
    if err != nil {
        return nil, err
    }

    query := r.db.Model(&consultation.ConsultationEntity{}).
        Where("check_in_at >= ? AND check_in_at <= ?", startDateTime, endDateTime)

    if clinicId != nil && *clinicId != "" {
        query = query.Where("clinic_id = ?", *clinicId)
    }
    if timePeriod != nil && *timePeriod != "" {
        query = query.Where("time_period = ?", *timePeriod)
    }
    if patientName != nil && *patientName != "" {
        query = query.Where("full_name ILIKE ?", "%"+*patientName+"%")
    }
    if patientId != nil && *patientId != "" {
        query = query.Where("patient_id = ?", *patientId)
    }
    if doctorId != nil && *doctorId != "" {
        query = query.Where("doctor_id = ?", *doctorId)
    }

    if err := query.Count(&totalCount).Error; err != nil {
        return nil, err
    }

    if err := query.Limit(limit).Offset(offset).Find(&consultations).Error; err != nil {
        return nil, err
    }

    var sharedConsultations []shared.ConsultationDetail
    for _, entity := range consultations {
        var acupunctureTreatment *shared.AcupunctureTreatment
        var medicineTreatment *shared.MedicineTreatment

        if entity.AcupunctureTreatment != nil {
            acupunctureTreatment = &shared.AcupunctureTreatment{
                ID: *entity.AcupunctureTreatment,
            }
        }

        if entity.MedicineTreatment != nil {
            medicineTreatment = &shared.MedicineTreatment{
                ID: *entity.MedicineTreatment,
            }
        }

        sharedConsultations = append(sharedConsultations, shared.ConsultationDetail{
            ID:                   entity.ID,
            Status:               shared.ConsultationStatus(entity.Status),
            Source:               shared.ConsultationSource(entity.Source),
            ConsultationNumber:   entity.ConsultationNumber,
            CheckInAt:            entity.CheckInAt,
            StartAt:              entity.StartAt,
            EndAt:                entity.EndAt,
            CheckOutAt:           entity.CheckOutAt,
            OnsiteCancelAt:       entity.OnsiteCancelAt,
            OnsiteCancelReason:   (*shared.OnsiteCancelReasonType)(entity.OnsiteCancelReason),
            IsFirstTimeVisit:     entity.IsFirstTimeVisit,
            PatientID:            entity.PatientID,
            TimeSlotID:           entity.TimeSlotID,
            AcupunctureTreatment: acupunctureTreatment,
            MedicineTreatment:    medicineTreatment,
        })
    }

    result := &shared.ConsultationQueryResult{
        Consultations: sharedConsultations,
        TotalCount:    int(totalCount),
    }

    return result, nil
}
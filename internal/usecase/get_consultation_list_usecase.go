package usecase

import (
	"clinic-pulse-go/internal/domain/consultation/interfaces/repositories"
	"clinic-pulse-go/internal/shared"
)

type GetConsultationListRequest struct {
	StartDate        string
	EndDate          string
	ClinicId         *string
	TimePeriod       *shared.TimePeriodType
	TotalDurationMin *int
	TotalDurationMax *int
	PatientName      *string
	PatientId        *string
	DoctorId         *string
	Page             int
	Limit            int
}

type GetConsultationListResponse struct {
	Data []shared.ConsultationDetail
	Pagination struct {
		Pages       []int
		TotalPage   int
		CurrentPage int
		Prev        int
		Next        int
	}
	TotalCounts int
}

type GetConsultationListUseCase struct {
	consultationRepo repositories.IConsultationRepository
}

func NewGetConsultationListUseCase(cr repositories.IConsultationRepository) *GetConsultationListUseCase {
	return &GetConsultationListUseCase{consultationRepo: cr}
}

func (uc *GetConsultationListUseCase) Execute(request GetConsultationListRequest) (*GetConsultationListResponse, error) {
    offset := (request.Page - 1) * request.Limit

    clinicId := request.ClinicId

    timePeriod := ""
    if request.TimePeriod != nil {
        timePeriod = string(*request.TimePeriod) 
    }

    totalDurationMin := request.TotalDurationMin
    totalDurationMax := request.TotalDurationMax

    patientName := request.PatientName
    patientId := request.PatientId
    doctorId := request.DoctorId

    result, err := uc.consultationRepo.FindByQuery(
        request.Limit, offset, request.StartDate, request.EndDate,
        clinicId, &timePeriod, totalDurationMin, totalDurationMax,
        patientName, patientId, doctorId,
    )

    if err != nil {
        return nil, err
    }

    response := &GetConsultationListResponse{
        Data:        result.Consultations,
        TotalCounts: result.TotalCount,
        Pagination: struct {
            Pages       []int
            TotalPage   int
            CurrentPage int
            Prev        int
            Next        int
        }{},
    }

    return response, nil
}
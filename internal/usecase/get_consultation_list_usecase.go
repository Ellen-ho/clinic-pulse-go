package usecase

import (
	"my-go-backend/internal/domain"
	"my-go-backend/internal/repository"
)

type GetConsultationListRequest struct {
	StartDate        string
	EndDate          string
	ClinicId         *string
	TimePeriod       *domain.TimePeriodType
	TotalDurationMin *int
	TotalDurationMax *int
	PatientName      *string
	PatientId        *string
	DoctorId         *string
	Page             int
	Limit            int
}

type GetConsultationListResponse struct {
	Data []domain.Consultation
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
	consultationRepo repository.IConsultationRepository
}

func NewGetConsultationListUseCase(cr repository.IConsultationRepository) *GetConsultationListUseCase {
	return &GetConsultationListUseCase{consultationRepo: cr}
}

func (uc *GetConsultationListUseCase) Execute(request GetConsultationListRequest) (*GetConsultationListResponse, error) {
	offset := (request.Page - 1) * request.Limit

	timePeriod := ""
	if request.TimePeriod != nil {
		timePeriod = string(*request.TimePeriod) 
	}

	clinicId := ""
	if request.ClinicId != nil {
		clinicId = *request.ClinicId
	}

	totalDurationMin := 0
	if request.TotalDurationMin != nil {
		totalDurationMin = *request.TotalDurationMin
	}

	totalDurationMax := 0
	if request.TotalDurationMax != nil {
		totalDurationMax = *request.TotalDurationMax
	}

	patientName := ""
	if request.PatientName != nil {
		patientName = *request.PatientName
	}

	patientId := ""
	if request.PatientId != nil {
		patientId = *request.PatientId
	}

	doctorId := ""
	if request.DoctorId != nil {
		doctorId = *request.DoctorId
	}

	consultations, totalCounts, err := uc.consultationRepo.FindByQuery(
		request.Limit, offset, request.StartDate, request.EndDate, clinicId,
		timePeriod, totalDurationMin, totalDurationMax,
		patientName, patientId, doctorId,
	)

	if err != nil {
		return nil, err
	}

	response := &GetConsultationListResponse{
		Data: consultations,
		TotalCounts: totalCounts,
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
package handler

import (
	"clinic-pulse-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
    "clinic-pulse-go/internal/shared"
)

type ConsultationHandler struct {
	getConsultationListUseCase *usecase.GetConsultationListUseCase
}

func NewConsultationHandler(uc *usecase.GetConsultationListUseCase) *ConsultationHandler {
	return &ConsultationHandler{getConsultationListUseCase: uc}
}

func (h *ConsultationHandler) GetConsultationList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))        
	page, _ := strconv.Atoi(c.Query("page", "1"))          
	startDate := c.Query("startDate", "")
	endDate := c.Query("endDate", "")
	clinicId := c.Query("clinicId", "")
	timePeriod := c.Query("timePeriod", "")
	totalDurationMin, _ := strconv.Atoi(c.Query("totalDurationMin", "0"))
	totalDurationMax, _ := strconv.Atoi(c.Query("totalDurationMax", "0"))
	patientName := c.Query("patientName", "")
	patientId := c.Query("patientId", "")
	doctorId := c.Query("doctorId", "")

	var timePeriodType *shared.TimePeriodType
	if timePeriod != "" {
		period := shared.TimePeriodType(timePeriod)
		timePeriodType = &period
	}

	req := usecase.GetConsultationListRequest{
		StartDate:        startDate,
		EndDate:          endDate,
		ClinicId:         &clinicId,
		TimePeriod:       timePeriodType, 
		TotalDurationMin: &totalDurationMin,
		TotalDurationMax: &totalDurationMax,
		PatientName:      &patientName,
		PatientId:        &patientId,
		DoctorId:         &doctorId,
		Page:             page,
		Limit:            limit,
	}

	response, err := h.getConsultationListUseCase.Execute(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response)
}
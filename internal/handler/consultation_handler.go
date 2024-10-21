package handler

import (
	"clinic-pulse-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ConsultationHandler struct {
	getConsultationListUseCase *usecase.GetConsultationListUseCase
}

func NewConsultationHandler(uc *usecase.GetConsultationListUseCase) *ConsultationHandler {
	return &ConsultationHandler{getConsultationListUseCase: uc}
}

func (h *ConsultationHandler) GetConsultationList(c *fiber.Ctx) error {
    
    var req usecase.GetConsultationListRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
    }

    response, err := h.getConsultationListUseCase.Execute(req)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(response)
}
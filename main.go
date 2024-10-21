package main

import (
	"clinic-pulse-go/internal/handler"
	"clinic-pulse-go/internal/repository"
	"clinic-pulse-go/internal/usecase"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	consultationRepo := repository.NewConsultationRepository()
	getConsultationListUC := usecase.NewGetConsultationListUseCase(consultationRepo)
	consultationHandler := handler.NewConsultationHandler(getConsultationListUC)

    app.Get("/consultations", consultationHandler.GetConsultationList)

	app.Listen(":3000")
}
package main

import (
	"my-go-backend/internal/handler"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/usecase"
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
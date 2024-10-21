package main

import (
	"clinic-pulse-go/internal/handler"
	"clinic-pulse-go/internal/repository"
	"clinic-pulse-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	consultationRepo := repository.NewConsultationRepository()
	getConsultationListUC := usecase.NewGetConsultationListUseCase(consultationRepo)
	consultationHandler := handler.NewConsultationHandler(getConsultationListUC)

	userRepo := repository.NewUserRepository(db) 
	userUsecase := usecase.NewUserUsecase(userRepo)

	app.Get("/consultations", consultationHandler.GetConsultationList)

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")  

		id, err := strconv.ParseUint(idParam, 10, 32)  
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		user, err := userUsecase.GetUser(int(id))  
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		return c.JSON(user)   
	})

	app.Listen(":3000")
}
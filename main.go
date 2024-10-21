package main

import (
	"clinic-pulse-go/internal/handler"
	"clinic-pulse-go/internal/repository"
	"clinic-pulse-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"                     
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB_NAME") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=disable"

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
		id := c.Params("id")  
	
		user, err := userUsecase.GetUser(id)  
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
	
		return c.JSON(user)   
	})

	app.Listen(":3000")
}
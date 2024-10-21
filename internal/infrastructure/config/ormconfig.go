package config

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetOrmConfig() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	isProd := os.Getenv("NODE_ENV") == "production"

	entitiesPath := "./src/infrastructure/entities/*.go"
	if isProd {
		entitiesPath = "./build/infrastructure/entities/*.go"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if isProd {
		db.Logger.LogMode(0)
	} else {
		db.Logger.LogMode(1)
	}

	return db
}

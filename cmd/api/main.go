package main

import (
	"log"
	"os"
	"tixeve-api/internal/database"
	"tixeve-api/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB:", err)
	}

	defer sqlDB.Close()

	app := fiber.New()

	routes.Main(app, db)

	port := os.Getenv("APP_PORT")

	log.Fatal(app.Listen(":" + port))

}

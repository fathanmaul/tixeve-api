package routes

import (
	"fmt"
	"os"
	"tixeve-api/internal/controllers"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Main(router fiber.Router, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	router.Route("/api/v1", func(v1 fiber.Router) {
		v1.Get("/", func(c fiber.Ctx) error {
			// get app_name from env with fallback
			appName := os.Getenv("APP_NAME")
			if appName == "" {
				appName = "TixEve API"
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": fmt.Sprintf("Welcome to %s!", appName),
			})
		})
		v1.Route("/users", func(users fiber.Router) {
			users.Get("/", userController.Index)
		})
	})
}

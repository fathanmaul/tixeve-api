package controllers

import (
	"tixeve-api/internal/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (controller *UserController) Index(c fiber.Ctx) error {
	var users []models.User

	if err := controller.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

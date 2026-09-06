package routes

import "github.com/gofiber/fiber/v3"

func Main(router fiber.Router) {
	router.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})
}

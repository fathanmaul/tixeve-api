package api

import (
	"log"
	"tixeve-api/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Route("/api/v1", routes.Main)

	log.Fatal(app.Listen(":3030"))
}

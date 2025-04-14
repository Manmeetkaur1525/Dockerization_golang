package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manmeetkaur1525/dockerization_golang/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	SetupRoutes(app)

	app.Listen(":3000")
}

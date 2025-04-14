package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manmeetkaur1525/dockerization_golang/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}

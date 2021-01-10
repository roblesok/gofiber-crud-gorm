package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/roblesok/gorm-fiber-crud/router"
)

func main() {
	port := ":3000"
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	router.BookRouter(api)

	log.Fatal(app.Listen(port))
}

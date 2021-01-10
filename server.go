package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/roblesok/gorm-fiber-crud/repository"
	"github.com/roblesok/gorm-fiber-crud/router"
	"github.com/roblesok/gorm-fiber-crud/service"
)

var (
	pgRepo      repository.BookRepository = repository.NewPgRepository()
	bookService service.BookService       = service.NewBookService(pgRepo)
)

func main() {
	port := ":3000"
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	router.BookRouter(api, bookService)

	log.Fatal(app.Listen(port))
}

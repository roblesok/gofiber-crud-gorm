package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/roblesok/gorm-fiber-crud/controller"
)

var (
	bookController controller.BookController = controller.NewBookController()
)

func main() {
	port := ":3000"
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This Works")
	})

	app.Get("/books", bookController.GetAllBooks)
	app.Get("/books/:id", bookController.GetBook)
	app.Post("/books", bookController.AddBook)
	app.Delete("/books/:id", bookController.DeleteBook)

	log.Fatal(app.Listen(port))
}

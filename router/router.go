package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesok/gorm-fiber-crud/controller"
)

var (
	bookController controller.BookController = controller.NewBookController()
)

func BookRouter(app fiber.Router) {
	app.Get("/books", bookController.GetAllBooks)
	app.Get("/books/:id", bookController.GetBook)
	app.Post("/books", bookController.AddBook)
	app.Delete("/books/:id", bookController.DeleteBook)
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesok/gorm-fiber-crud/controller"
	"github.com/roblesok/gorm-fiber-crud/service"
)

var (
	bookController controller.BookController = controller.NewBookController()
)

func BookRouter(app fiber.Router, service service.BookService) {
	app.Get("/books", bookController.GetAllBooks(service))
	app.Get("/books/:id", bookController.GetBook(service))
	app.Post("/books", bookController.AddBook(service))
	app.Delete("/books/:id", bookController.DeleteBook(service))
}

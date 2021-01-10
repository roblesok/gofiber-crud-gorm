package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesok/gorm-fiber-crud/service"
)

type controller struct{}

type BookController interface {
	GetAllBooks(svc service.BookService) fiber.Handler
	GetBook(svc service.BookService) fiber.Handler
	AddBook(svc service.BookService) fiber.Handler
	DeleteBook(svc service.BookService) fiber.Handler
}

func NewBookController() BookController {
	return &controller{}
}

func (*controller) GetAllBooks(svc service.BookService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fetched, err := svc.FetchAll()
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"err": err,
			})
		}
		return ctx.JSON(&fiber.Map{
			"books": fetched,
		})
	}
}

func (*controller) GetBook(svc service.BookService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("TODO GET ONE")
	}
}

func (*controller) AddBook(svc service.BookService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("TODO ADD")
	}
}

func (*controller) DeleteBook(svc service.BookService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("TODO DELETE")
	}

}

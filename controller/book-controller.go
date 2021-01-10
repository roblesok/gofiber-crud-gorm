package controller

import "github.com/gofiber/fiber/v2"

type controller struct{}

type BookController interface {
	GetAllBooks(c *fiber.Ctx) error
	GetBook(c *fiber.Ctx) error
	AddBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}

func NewBookController() BookController {
	return &controller{}
}

func (*controller) GetAllBooks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ALL BOOKS",
	})
}

func (*controller) GetBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ONE BOOK",
	})
}

func (*controller) AddBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ADD BOOK",
	})
}

func (*controller) DeleteBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DELETE BOOK",
	})
}

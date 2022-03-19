package handler

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Books fetched sucessfully")
}

func GetBook(c *fiber.Ctx) {
	c.SendString("Book fetched sucessfully")
}
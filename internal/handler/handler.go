package handler

import "github.com/gofiber/fiber/v2"

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("hello world")
}

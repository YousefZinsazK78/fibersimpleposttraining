package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/database"
)

type handler struct {
	userer database.Userer
	poster database.Poster
}

func NewHandler(userer database.Userer, poster database.Poster) handler {
	return handler{
		userer: userer,
		poster: poster,
	}
}

func (h handler) Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("hello world")
}

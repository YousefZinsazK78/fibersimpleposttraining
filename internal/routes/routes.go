package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/handler"
)

func Run(port string) {
	var (
		app    = fiber.New()
		hndler = handler.NewHandler()
	)

	app.Get("/hello", hndler.Hello)

	app.Listen(port)
}

package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	///set port in flag of program
	port := flag.String("default port", ":8000", "you can set your custom font")
	flag.Parse()

	//create new instance of fiber
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": "hello world"})
	})

	//listen to fiber app
	app.Listen(*port)
}

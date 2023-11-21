package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	code := fiber.StatusInternalServerError
	message := "internal server error!!"
	if e, ok := err.(models.CustomErrorBlog); ok {
		code = e.Code
		message = e.Message
	}

	return ctx.Status(code).JSON(fiber.Map{"error_msg": message})
}

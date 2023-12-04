package models

import (
	"github.com/gofiber/fiber/v2"
)

type CustomErrorBlog struct {
	Code    int
	Message string
}

func (c CustomErrorBlog) Error() string {
	return c.Message
}

func InternalServerError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusInternalServerError,
		Message: "Internal Server Error! 🤨🤨",
	}
}

func UnAuthorizedError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusUnauthorized,
		Message: "UnAuthorized! 🤨🤨",
	}
}

func TokenExpiredError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusUnauthorized,
		Message: "Token Expired! 🤨🤨",
	}
}

func CantConvertError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusInternalServerError,
		Message: "Can't Convert ID Params! 🥱😫",
	}
}

func NotFoundError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusNotFound,
		Message: "Not Found Error! 😲🤯",
	}
}

func InvalidFieldError() error {
	return CustomErrorBlog{
		Code:    fiber.StatusInternalServerError,
		Message: "Invalid Field Pass Error! 😑😐",
	}
}

func NewCustomBlogError(code int, message string) error {
	return CustomErrorBlog{
		Code:    code,
		Message: message,
	}
}

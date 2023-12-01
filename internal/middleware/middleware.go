package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/helper"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return errors.New("unauthorized!")
		}
		tStr := strings.Split(token, "Bearer ")
		fmt.Println(tStr)
		if helper.VaildateJwtToken(tStr[1]) != nil {
			return errors.New("unauthorized!")
		}
		username, err := helper.GetUsernameJwtToken(tStr[1])
		if err != nil {
			return errors.New("unauthorized!")
		}
		fmt.Println(username)

		return c.Next()
	}
}

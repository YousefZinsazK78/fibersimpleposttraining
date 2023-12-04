package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/helper"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return models.UnAuthorizedError()
		}
		if !strings.Contains(token, "Bearer ") {
			return models.UnAuthorizedError()
		}
		tStr := strings.Split(token, "Bearer ")
		if helper.VaildateJwtToken(tStr[1]) != nil {
			return models.TokenExpiredError()
		}
		username, err := helper.GetUsernameJwtToken(tStr[1])
		if err != nil {
			return models.UnAuthorizedError()
		}
		fmt.Println(username)
		return c.Next()
	}
}

package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/helper"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func (h handler) UserInsert(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	var userInsertModel *models.UserInsertParams
	if err := c.BodyParser(userInsertModel); err != nil {
		return err
	}

	err := h.userer.Insert(timeoutContext, *userInsertModel)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": "user registered successfully",
	})
}

func (h handler) UserLogin(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	var userLoginModel *models.UserLoginParams
	if err := c.BodyParser(userLoginModel); err != nil {
		return err
	}

	user, err := h.userer.GetByUsername(timeoutContext, userLoginModel.Username)
	if err != nil {
		return err
	}

	if helper.ComparePassword(user.Password, userLoginModel.Password) != nil {
		return models.NewCustomBlogError(fiber.StatusNotFound, "invalid credintials")
	}

	tokenstring := helper.GenerateJwtToken(user.Username)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result":   "user login successfully",
		"jwtToken": tokenstring,
	})
}

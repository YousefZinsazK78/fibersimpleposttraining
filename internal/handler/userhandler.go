package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/helper"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func (h handler) UserInsert(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	var userInsertModel models.UserInsertParams
	if err := c.BodyParser(&userInsertModel); err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(userInsertModel.Password)
	if err != nil {
		return models.NewCustomBlogError(fiber.StatusInternalServerError, err.Error())
	}

	userInsertModel.Password = hashedPassword

	err = h.userer.Insert(timeoutContext, userInsertModel)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": "user registered successfully",
	})
}

func (h handler) UserLogin(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	var userLoginModel models.UserLoginParams
	if err := c.BodyParser(&userLoginModel); err != nil {
		return err
	}

	user, err := h.userer.GetByUsername(timeoutContext, userLoginModel.Username)
	if err != nil {
		return err
	}

	if !helper.ComparePassword(user.Password, userLoginModel.Password) {
		return models.NewCustomBlogError(fiber.StatusInternalServerError, "invalid credintials")
	}

	tokenstring := helper.GenerateJwtToken(user.Username)

	c.Request().Header.Add("Authorization", "Bearer "+tokenstring)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result":   "user login successfully",
		"jwtToken": tokenstring,
	})
}

func (h handler) GetUsers(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	users, err := h.userer.GetUsers(timeoutContext)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": users,
	})
}

func (h handler) GetUserByID(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.CantConvertError()
	}

	user, err := h.userer.GetUserByID(timeoutContext, id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": user,
	})
}

func (h handler) GetUserByEmail(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	user, err := h.userer.GetUserByEmail(timeoutContext, c.Params("email"))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": user,
	})
}

func (h handler) GetByUsername(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Millisecond*100)
	defer cancel()

	user, err := h.userer.GetByUsername(timeoutContext, c.Params("username"))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": user,
	})
}

package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func (h handler) PostInsert(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	var postInsertModel models.PostInsertParams
	if err := c.BodyParser(&postInsertModel); err != nil {
		return err
	}

	err := h.poster.Insert(timeoutContext, postInsertModel)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": "your post insert successfully",
	})
}

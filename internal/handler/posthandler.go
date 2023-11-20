package handler

import (
	"context"
	"strconv"
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

func (h handler) GetPosts(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	postList, err := h.poster.GetPosts(timeoutContext)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": postList,
	})
}

func (h handler) GetPostByID(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't convert id (id is still string 😥😥",
		})
	}

	post, err := h.poster.GetPostByID(timeoutContext, id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": post,
	})
}

func (h handler) GetPostByTitle(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()

	posts, err := h.poster.GetPostByTitle(timeoutContext, c.Params("title"))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": posts,
	})
}

package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

func (h handler) PostInsert(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	var postInsertModel models.PostInsertParams
	if err := c.BodyParser(&postInsertModel); err != nil {
		return models.InvalidFieldError()
	}

	err := h.poster.Insert(timeoutContext, postInsertModel)
	if err != nil {
		return models.InternalServerError()
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": "your post insert successfully",
	})
}

func (h handler) GetPosts(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	postList, err := h.poster.GetPosts(timeoutContext)
	if err != nil {
		return models.InternalServerError()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": postList,
	})
}

func (h handler) GetPostByID(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.CantConvertError()
	}

	post, err := h.poster.GetPostByID(timeoutContext, id)
	if err != nil {
		return models.InternalServerError()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": post,
	})
}

func (h handler) GetPostByTitle(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	posts, err := h.poster.GetPostByTitle(timeoutContext, c.Params("title"))
	if err != nil {
		return models.CantConvertError()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": posts,
	})
}

func (h handler) PutPost(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	//body parser
	var postmodel models.PostUpdateParams
	if err := c.BodyParser(&postmodel); err != nil {
		return models.InvalidFieldError()
	}

	post, err := h.poster.Update(timeoutContext, postmodel)
	if err != nil {
		return models.InternalServerError()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": post,
	})
}

func (h handler) DeletePost(c *fiber.Ctx) error {
	timeoutContext, cancel := context.WithTimeout(c.Context(), time.Second*1)
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.CantConvertError()
	}

	err = h.poster.Delete(timeoutContext, id)
	if err != nil {
		return models.InternalServerError()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": fmt.Sprintf("%d delete successfuly", id),
	})
}

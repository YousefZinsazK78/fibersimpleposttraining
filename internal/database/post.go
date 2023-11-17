package database

import (
	"context"

	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Poster interface {
	Insert(context.Context, models.PostInsertParams) error
}

type post struct{}

func NewPostDB() post {
	return post{}
}

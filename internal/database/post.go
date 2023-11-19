package database

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Poster interface {
	Insert(context.Context, models.PostInsertParams) error
}

type post struct {
	database
}

func NewPostDB(db database) post {
	return post{
		database: db,
	}
}

func (p post) Insert(ctx context.Context, postModel models.PostInsertParams) error {
	sqlB := sqlbuilder.NewInsertBuilder()
	sqlB.InsertInto("post_tbl")
	sqlB.Cols("title", "content")
	sqlB.Values(postModel.Title, postModel.Content)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := p.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

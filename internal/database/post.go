package database

import (
	"context"

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
	//todo : write sql query with sqlbuilder
	//todo : use db

	// p.db.QueryContext(ctx, sql, args)
	return nil
}

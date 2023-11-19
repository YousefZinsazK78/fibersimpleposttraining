package database

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Poster interface {
	Insert(context.Context, models.PostInsertParams) error
	GetPosts(context.Context) ([]models.Post, error)
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

	err := p.db.QueryRowContext(ctx, sql, args...)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (p post) GetPosts(ctx context.Context) ([]models.Post, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("post_tbl")
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := p.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

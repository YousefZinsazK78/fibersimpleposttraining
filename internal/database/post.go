package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Poster interface {
	Insert(context.Context, models.PostInsertParams) error
	GetPosts(context.Context) ([]models.Post, error)
	GetPostByID(context.Context, int) (*models.Post, error)
	GetPostByTitle(context.Context, string) ([]models.Post, error)
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

func (p post) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("post_tbl")
	sqlB.Where(
		sqlB.Equal("post_id", id),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := p.db.QueryRowContext(ctx, sql, args...)

	var post models.Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return nil, err
	}

	return &post, nil
}

func (p post) GetPostByTitle(ctx context.Context, title string) ([]models.Post, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("post_tbl")
	sqlB.Where(
		sqlB.Like("title", fmt.Sprintf("%%%s%%", strings.ReplaceAll(title, "%20", " "))),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := p.db.QueryContext(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
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

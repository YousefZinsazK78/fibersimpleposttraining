package database

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Userer interface {
	Insert(context.Context, models.UserInsertParams) error
	GetByUsername(context.Context, string) (*models.User, error)
}

type userDB struct {
	database
}

func NewUserDB(db database) userDB {
	return userDB{
		database: db,
	}
}

func (u userDB) Insert(ctx context.Context, userModel models.UserInsertParams) error {
	//todo : use sql builder
	sqlB := sqlbuilder.NewInsertBuilder()
	sqlB.InsertInto("user_tbl")
	sqlB.Cols("username", "HashPassword", "email")
	sqlB.Values(userModel.Username, userModel.Password, userModel.Email)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := u.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u userDB) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("user_tbl")
	sqlB.Where(
		sqlB.Equal("username", username),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := u.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

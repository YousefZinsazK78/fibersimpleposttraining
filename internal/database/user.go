package database

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Userer interface {
	Insert(context.Context, models.UserInsertParams) error
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
	sqlB.InsertInto("users")
	sqlB.Cols("username", "password", "email")
	sqlB.Values(userModel.Username, userModel.Password, userModel.Email)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := u.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

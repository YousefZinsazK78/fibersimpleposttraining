package database

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
)

type Userer interface {
	Insert(context.Context, models.UserInsertParams) error
	GetByUsername(context.Context, string) (*models.User, error)
	GetUsers(context.Context) ([]models.User, error)
	GetUserByID(context.Context, int) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	Update(context.Context, models.UserUpdateParams) (*models.User, error)
	Delete(context.Context, int) error
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
	fmt.Println(userModel)
	sqlB := sqlbuilder.NewInsertBuilder()
	sqlB.InsertInto("user_tbl")
	sqlB.Cols("Username", "Email", "HashPassword")
	sqlB.Values(userModel.Username, userModel.Email, userModel.Password)
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
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userDB) GetUsers(ctx context.Context) ([]models.User, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("user_tbl")
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := u.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u userDB) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("user_tbl")
	sqlB.Where(
		sqlB.Equal("user_id", id),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := u.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userDB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	sqlB := sqlbuilder.NewSelectBuilder()
	sqlB.Select("*")
	sqlB.From("user_tbl")
	sqlB.Where(
		sqlB.Equal("email", email),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := u.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userDB) Update(ctx context.Context, userModel models.UserUpdateParams) (*models.User, error) {
	sqlB := sqlbuilder.NewUpdateBuilder()
	sqlB.Update("user_tbl")
	sqlB.Set(
		sqlB.Assign("username", userModel.Username),
		"updatedat = CURRENT_TIMESTAMP",
	)
	sqlB.Where(
		sqlB.Equal("user_id", userModel.ID),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := u.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user, err := u.GetUserByID(ctx, userModel.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userDB) Delete(ctx context.Context, id int) error {
	sqlB := sqlbuilder.NewDeleteBuilder()
	sqlB.DeleteFrom("user_tbl")
	sqlB.Where(
		sqlB.Equal("user_id", id),
	)
	sql, args := sqlB.BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := u.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

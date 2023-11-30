package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SuperSecretPassword = "supersimplesecretkeypasswordin2023year21centuryh1h1@ZzZ#"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type UserInsertParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type PostInsertParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdateParams struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CustomUserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

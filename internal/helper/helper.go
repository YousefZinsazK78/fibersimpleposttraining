package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func GenerateJwtToken(username string) string {
	customClaims := models.CustomUserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "GenerateJwtToken",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tstring, err := token.SignedString([]byte(models.SuperSecretPassword))
	if err != nil {
		return ""
	}
	return tstring
}

func VaildateJwtToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(models.SuperSecretPassword), nil
	})

	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*models.CustomUserClaims)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return errors.New("token expired!")
	}
	return nil
}

func GetUsernameJwtToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(models.SuperSecretPassword), nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*models.CustomUserClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}
	return claims.Username, nil
}

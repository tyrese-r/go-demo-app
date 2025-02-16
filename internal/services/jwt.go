package services

import (
	"github.com/golang-jwt/jwt/v5"
	"go-demo-app/internal/utils/secrets"
	"time"
)

var jwtSecret = []byte(secrets.GetFromEnv("JWT_SECRET", "default-secret"))

func CreateJWTToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"roles":    []string{"user"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

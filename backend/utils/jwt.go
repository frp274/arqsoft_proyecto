package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuracion = time.Hour * 24
	jwtSecret   = "jwtSecret"
)

func GenerateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(jwtDuracion)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "backend",
		Subject:    "auth",
		ID:        fmt.Sprintf("%d", userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("error generating token %w", err)
	}

	return tokenString, nil
}

package utils

import (
	"fmt"
	"time"
	"log"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Printf("error al generar el token %v", err)
		return "", fmt.Errorf("error generating token %w", err)
	}

	return tokenString, nil
}

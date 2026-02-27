package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecret = "jwtSecret"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Es_admin bool `json:"es_admin"`
}

func ParseJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

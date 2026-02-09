package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret_key_123")

type MyClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	claims := MyClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)

	return token, err
}

func ValidateToken(tkn string) (*MyClaims, error) {
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tkn, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("abcde")

type CustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(email, role string) (string, error) {
	accessTTL := time.Minute * 20
	accessClaims := &CustomClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretkey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateToken(str string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(str, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Method")

		}
		return secretkey, nil
	})
	if err != nil || !token.Valid{
		return nil, errors.New("invalid token")

	}
	return claims,nil
}

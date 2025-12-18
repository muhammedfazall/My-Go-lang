package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

type CustomClaims struct{
	Email    string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokens(email, role string) (string, string, error) {
	accessTime := time.Minute * 15 
	accessClaims := &CustomClaims{
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTime)),
		},
	}

	accessToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256,accessClaims).SignedString(secretKey)
	if err != nil{
		return "","",err
	}

	refreshTime := time.Hour * 24 * 7
	refreshClaims := &CustomClaims{
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTime)),
		},
	}

	refreshToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims).SignedString(secretKey)
	if err != nil{
		return "","",err
	}

	return accessToken,refreshToken,nil
}

func ValidateToken(tokenString string) (*CustomClaims,error) {

	claims := &CustomClaims{}

	token,err := jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (any, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,errors.New("invalid method")
		}
		return secretKey,nil
	})

	if err != nil || !token.Valid {
		return nil,errors.New("invalid Token")
	}
	return claims,nil
}
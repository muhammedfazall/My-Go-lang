package myjwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")


func GenerateToken(role,email string,ttl time.Duration)(string, error){
	claims := jwt.MapClaims{
		"email" : email,
		"role": role,
		"exp" : time.Now().Add(ttl).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString(secretKey)
}

func ValidateToken(tokenstr string)(jwt.MapClaims,error){
	token,err := jwt.Parse(tokenstr,func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,errors.New("unexp sign in method")
		}
		return 	secretKey,nil
	})

	if err != nil || !token.Valid{
		return nil, errors.New("invalid token")
	}

	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil,errors.New("invalid claims")
	}

	exp,ok := claims["exp"].(float64)
	if !ok{
		return nil,errors.New("inalid exp")
	}
	if time.Now().Unix() > int64(exp){
		return nil,errors.New("token expired")
	}
	return claims,nil
}

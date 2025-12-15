package jwtoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("secret")

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// func GenerateToken(email string, ttl time.Duration) (string, error) {
// 	claims := jwt.MapClaims{
// 		"email": email,
// 		"exp":   time.Now().Add(ttl).Unix(),
// 	}

// 	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretkey)
// }

func GenerateAccessAndRefreshTokens(userID uint, email, role string) (accessToken string, refreshToken string, err error) {
	accessTTL := time.Minute * 15
	accessClaims := &CustomClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretkey)
	if err != nil {
		return
	}

	refreshTTL := time.Hour * 24 * 7
	refreshClaims := &CustomClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTTL)),
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretkey)
	if err != nil {
		return
	}

	return accessToken, refreshToken, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid method")
		}
		return secretkey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil

	// exp := claims.ExpiresAt

	// if time.Now().Unix() > int64(exp) {
	// 	return nil, errors.New("token expired")
	// }

}

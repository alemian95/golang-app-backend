package auth

import (
	"fmt"
	"time"

	"alessandromian.dev/golang-app/app/models/user_model"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthClaims struct {
	UserId uint  `json:"user_id"`
	Exp    int64 `json:"exp"`
	jwt.StandardClaims
}

func GenerateToken(user_id uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	claims := &AuthClaims{
		UserId: user_id,
		Exp:    expirationTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (AuthClaims, error) {

	claims := &AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return *claims, fmt.Errorf("token is either expired or not active yet")
			}
		}
		return *claims, err
	}

	if !token.Valid {
		return *claims, fmt.Errorf("invalid token")
	}

	return *claims, nil
}

func GetUserBySession(session string) *user_model.User {
	claims, err := VerifyToken(session)

	if err != nil {
		return nil
	}

	user, err := user_model.Find(uint64(claims.UserId))

	if err != nil {
		return nil
	}

	return user
}

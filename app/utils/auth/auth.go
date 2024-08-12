package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"golang-app/app/models/user_model"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email"`
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type AuthClaims struct {
	// UserId uint  `json:"user_id"`
	Payload map[string]any `json:"payload"`
	Exp     int64          `json:"exp"`
	jwt.StandardClaims
}

// func GenerateToken(user_id uint) (string, error) {
func GenerateToken(payload map[string]any) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	claims := &AuthClaims{
		Payload: payload,
		Exp:     expirationTime,
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

func GetUserBySession(session string) (*user_model.User, error) {
	claims, err := VerifyToken(session)

	if err != nil {
		return nil, err
	}

	user, err := user_model.Find(uint64(claims.Payload["user_id"].(float64)))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GenerateRandomToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package validator

import (
	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/utils/auth"
	"github.com/go-playground/validator/v10"
)

func ValidateUser(user *user_model.User) error {
	return validator.New().Struct(user)
}

func ValidateLoginRequest(login *auth.LoginRequest) error {
	return validator.New().Struct(login)
}

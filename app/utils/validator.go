package util

import (
	"alessandromian.dev/golang-app/app/models/user_model"
	"github.com/go-playground/validator/v10"
)

func Validate(user *user_model.User) error {
	return validator.New().Struct(user)
}

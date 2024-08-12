package forgot_password_token_model

import (
	"golang-app/app/models/database"
)

type ForgotPasswordToken struct {
	UserId uint `gorm:"unique"`
	Token  string
}

func All() []ForgotPasswordToken {
	var tokens []ForgotPasswordToken
	database.Conn().Find(&tokens)
	return tokens
}

func FindByUser(user_id uint) (*ForgotPasswordToken, error) {
	var token ForgotPasswordToken
	err := database.Conn().Where("user_id = ?", user_id).First(&token).Error
	return &token, err
}

func FindByToken(token_str string) (*ForgotPasswordToken, error) {
	var token ForgotPasswordToken
	err := database.Conn().Where("token = ?", token_str).First(&token).Error
	return &token, err
}

func (t *ForgotPasswordToken) Create() error {
	return database.Conn().Create(t).Error
}

func (t *ForgotPasswordToken) Delete() error {
	return database.Conn().Delete("user = ?", t.UserId).Error
}

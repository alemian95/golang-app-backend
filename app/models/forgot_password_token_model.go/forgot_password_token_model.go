package forgot_password_token_model

import (
	"golang-app/app/models/database"
	"golang-app/app/models/user_model"
)

type ForgotPasswordToken struct {
	UserId uint            `gorm:"primarykey"`
	User   user_model.User `gorm:"foreignKey:UserId"`
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

func (t *ForgotPasswordToken) Save() error {
	return database.Conn().Save(&t).Error
}

func (t *ForgotPasswordToken) Delete() error {
	return database.Conn().Delete(&t).Error
}

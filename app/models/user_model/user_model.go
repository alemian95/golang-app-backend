package user_model

import (
	"golang-app/app/models/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func All() []User {
	var users []User
	database.Conn().Find(&users)
	return users
}

func Find(id uint64) (*User, error) {
	var user User
	err := database.Conn().First(&user, id).Error
	return &user, err
}

func FindByEmail(email string) (*User, error) {
	var user User
	err := database.Conn().Where("email =?", email).First(&user).Error
	return &user, err
}

func CheckIfEmailExists(email string) bool {
	var user User
	err := database.Conn().Where("email =?", email).First(&user).Error
	if err != nil {
		return false
	}
	if user.Email == "" {
		return false
	} else {
		return true
	}
}

func (u *User) Read() error {
	return database.Conn().First(&u, u.ID).Error
}

func (u *User) Create() error {
	return database.Conn().Create(u).Error
}

func (u *User) Update() error {
	return database.Conn().Save(&u).Error
}

func (u *User) Delete() error {
	return database.Conn().Delete(&u, u.ID).Error
}

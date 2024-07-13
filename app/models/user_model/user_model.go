package user_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func All(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func Find(db *gorm.DB, id uint64) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	return &user, err
}

func FindByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email =?", email).First(&user).Error
	return &user, err
}

func (u *User) Read(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(&u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(&u, u.ID).Error
}

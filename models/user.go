package models

import (
	"card-game/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Email  string `gorm:"unique"`
	Status string

	Wallet []Wallet
}

func CreateUser(user *User) (*User, error) {
	db := database.DBConn

	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func UpdateUser(user User, id uint64) error {
	db := database.DBConn

	user.ID = uint(id)

	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

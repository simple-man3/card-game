package services

import (
	"card-game/database"
	"card-game/models"
)

func CreateUser(user *models.User) error {
	db := database.DBConn

	res := db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func Update(user models.User) error {
	db := database.DBConn

	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ExistUser(user models.User) bool {
	db := database.DBConn

	return db.Where(user).Find(&user).RowsAffected != 0
}

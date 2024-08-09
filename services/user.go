package services

import (
	"card-game/database"
	"card-game/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	db := database.DBConn

	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func Update(user models.User) error {
	db := database.DBConn

	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

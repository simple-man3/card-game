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

func UpdateUser(user *models.User, id uint) error {
	db := database.DBConn

	user.ID = id
	result := db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserById(id uint) (*models.User, error) {
	db := database.DBConn
	user := &models.User{ID: id}

	result := db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func DeleteUser(user models.User) error {
	db := database.DBConn

	result := db.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ExistUser(user models.User) bool {
	db := database.DBConn

	return db.Where(user).Find(&user).RowsAffected != 0
}

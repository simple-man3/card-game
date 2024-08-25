package services

import (
	"card-game/database"
	"card-game/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	db := database.DBConn

	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword
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

func GetUser(user *models.User, relations []string) error {
	query := database.DBConn

	if len(relations) > 0 {
		for _, relation := range relations {
			query = query.Preload(relation)
		}
	}

	result := query.Where(user).First(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

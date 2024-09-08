package services

import (
	"card-game/database"
	"card-game/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: database.DBConn,
	}
}

func (us UserService) CreateUser(user *models.User) error {
	hashPassword, err := us.hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword
	res := us.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (us UserService) UpdateUser(user *models.User, id uint) error {
	user.ID = id
	result := us.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (us UserService) GetUserById(id uint) (*models.User, error) {
	user := &models.User{ID: id}

	result := us.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (us UserService) DeleteUser(user models.User) error {
	result := us.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (us UserService) ExistUser(user models.User) bool {
	return us.db.Where(user).Find(&user).RowsAffected != 0
}

func (us UserService) GetUser(user *models.User, relations []string) error {
	query := us.db

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

func (us UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

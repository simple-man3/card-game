package services

import (
	"card-game/database"
	"card-game/models"
)

func CreateWallet(wallet *models.Wallet) error {
	db := database.DBConn

	res := db.Create(&wallet)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func PutMoney(amount float64, wallet *models.Wallet) error {
	db := database.DBConn

	wallet.Balance += amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func TakeMoney(amount float64, wallet *models.Wallet) error {
	db := database.DBConn

	wallet.Balance -= amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func GetWalletById(id uint, relations []string) (*models.Wallet, error) {
	query := database.DBConn
	wallet := models.Wallet{ID: id}

	if len(relations) > 0 {
		for _, relation := range relations {
			query = query.Preload(relation)
		}
	}

	result := query.First(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}

	return &wallet, nil
}

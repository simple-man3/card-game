package services

import (
	"card-game/database"
	"card-game/models"
)

func CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	db := database.DBConn

	res := db.Create(wallet)

	if res.Error != nil {
		return nil, res.Error
	}

	return wallet, nil
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

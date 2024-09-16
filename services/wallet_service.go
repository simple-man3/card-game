package services

import (
	"card-game/database"
	"card-game/models"
)

type WalletService struct{}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (ws WalletService) CreateWallet(wallet *models.Wallet) error {
	db := database.DBConn

	res := db.Create(&wallet)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (ws WalletService) PutMoney(amount float64, wallet *models.Wallet) error {
	db := database.DBConn

	wallet.Balance += amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (ws WalletService) TakeMoney(amount float64, wallet *models.Wallet) error {
	db := database.DBConn

	wallet.Balance -= amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (ws WalletService) GetWalletById(id, userId uint, relations []string) (*models.Wallet, error) {
	query := database.DBConn
	wallet := models.Wallet{ID: id, UserId: userId}

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

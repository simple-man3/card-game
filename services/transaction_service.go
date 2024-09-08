package services

import (
	"card-game/database"
	"card-game/models"
)

type TransactionService struct {
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (ts TransactionService) CreateTransaction(trx *models.Transaction) (*models.Transaction, error) {
	db := database.DBConn

	res := db.Create(&trx)
	if res.Error != nil {
		return nil, res.Error
	}

	return trx, nil
}

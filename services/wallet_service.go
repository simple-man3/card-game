package services

import (
	"card-game/consts"
	"card-game/database"
	"card-game/models"
	"gorm.io/gorm"
)

type WalletService struct {
	db         *gorm.DB
	trxService *TransactionService
}

func NewWalletService() *WalletService {
	return &WalletService{
		db:         database.DBConn,
		trxService: NewTransactionService(),
	}
}

func (ws WalletService) CreateWallet(wallet *models.Wallet) error {
	db := database.DBConn

	res := db.Create(&wallet)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (ws WalletService) PutMoney(amount float64) error {
	ws.db.Preload("Wallet").First(&models.AuthUser)

	models.AuthUser.Wallet.Balance += amount

	tx := ws.db.Begin()
	transactionDto := ws.fillTransactionToCreate(
		models.AuthUser.Wallet.ID,
		amount,
		consts.PutMoney,
	)
	if _, err := ws.trxService.Create(&transactionDto); err != nil {
		tx.Rollback()
		return err
	}

	if err := ws.Update(models.AuthUser.Wallet); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (ws WalletService) fillTransactionToCreate(
	walletId uint,
	sum float64,
	action consts.TrxAction,
) models.Transaction {
	return models.Transaction{
		WalletId: walletId,
		Sum:      sum,
		Action:   action,
	}
}

func (ws WalletService) Update(wallet *models.Wallet) error {
	res := ws.db.Save(&wallet)
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

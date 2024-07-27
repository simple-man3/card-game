package repositories

import "card-game/models"

type (
	IUserRepository interface {
		Create(user models.User) (models.User, error)
		Update(user models.User) error
		Delete(id uint64) error
		FinById(id uint64) (models.User, error)
	}
	IWalletRepository interface {
		Create(wallet models.Wallet) (models.Wallet, error)
		PutMoney(id uint64, money float64) error
		WithdrawMoney(id uint64, money float64) error
	}
	ITransactionRepository interface {
		Create(transaction models.Transaction) (models.Transaction, error)
	}
)

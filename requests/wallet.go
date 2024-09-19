package requests

type (
	CreateWalletRequest struct {
		UserId  uint    `validate:"required,user-exist" json:"user_id"`
		Balance float64 `validate:"required,min=1" json:"balance"`
	}
	PutMoneyRequest struct {
		Amount float64 `validate:"required,min=1" json:"amount"`
	}
)

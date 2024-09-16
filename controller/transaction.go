package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
}

func NewTransactionController() *TransactionController {
	return &TransactionController{}
}

func (tc TransactionController) PutMoney(c *fiber.Ctx) error {
	fmt.Println(12121)
	return nil
}

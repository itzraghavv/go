package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {

	if t.transactionType == transactionWithdrawal {
		if t.amount > c.balance {
			return errors.New("insufficient funds")
		}
		c.balance -= t.amount
	}
	if t.transactionType == transactionDeposit {
		c.balance += t.amount
	}

	if t.transactionType == "unknown" {
		return errors.New("unknown transaction type")
	}

	return nil
}

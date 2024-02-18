package domain

import (
	"time"
)

type Account struct {
	ID        AccountID
	balance   float64
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func NewAccount(balance float64) *Account {
	return &Account{
		balance:   balance,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (account *Account) validate() error {
	if account.balance < 0 {
		return ErrNegativeBalance
	}
	if account.createdAt.IsZero() {
		return ErrCreatedAtCannotBeNull
	}
	if account.updatedAt.IsZero() {
		return ErrUpdatedAtCannotBeNull
	}
	return nil
}

func (account *Account) Deposit(amount float64) error {
	account.balance += amount
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if account.balance < amount {
		return ErrInsufficientFunds
	}
	account.balance -= amount
	return nil
}

func (account *Account) Close() error {
	if account.balance != 0 {
		return ErrBalanceMustBeZeroToCloseAccount
	}

	account.deletedAt = time.Now()
	return nil
}

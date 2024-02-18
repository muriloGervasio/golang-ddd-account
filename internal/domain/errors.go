package domain

import "errors"

var (
	ErrNegativeBalance                 = errors.New("balance cannot be negative")
	ErrCreatedAtCannotBeNull           = errors.New("createdAt cannot be null")
	ErrUpdatedAtCannotBeNull           = errors.New("updatedAt cannot be null")
	ErrInsufficientFunds               = errors.New("insufficient funds")
	ErrBalanceMustBeZeroToCloseAccount = errors.New("balance must be zero to close account")
)

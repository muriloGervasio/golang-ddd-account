package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount(0)
	require.NotNil(t, account)
	require.Equal(t, float64(0), account.balance)
	require.False(t, account.createdAt.IsZero())
	require.False(t, account.updatedAt.IsZero())
}

func TestAccount_Deposit(t *testing.T) {
	account := NewAccount(0)
	err := account.Deposit(100)
	require.Nil(t, err)
	require.Equal(t, float64(100), account.balance)
}

func TestAccount_Withdraw(t *testing.T) {
	account := NewAccount(100)
	print(account.balance)
	err := account.Withdraw(50)
	require.Nil(t, err)
	require.Equal(t, float64(50), account.balance)
}

func TestAccount_Withdraw_InsufficientFunds(t *testing.T) {
	account := NewAccount(100)
	err := account.Withdraw(200)
	require.Equal(t, ErrInsufficientFunds, err)
}

func TestAccount_Close(t *testing.T) {
	account := NewAccount(0)
	err := account.Close()
	require.Nil(t, err)
	require.False(t, account.deletedAt.IsZero())
}

func TestAccount_Close_BalanceNotZero(t *testing.T) {
	account := NewAccount(100)
	err := account.Close()
	require.Equal(t, ErrBalanceMustBeZeroToCloseAccount, err)
}

func TestAccount_Validate(t *testing.T) {
	account := NewAccount(0)
	err := account.validate()
	require.Nil(t, err)
}

func TestAccount_Validate_NegativeBalance(t *testing.T) {
	account := NewAccount(-1)
	err := account.validate()
	require.Equal(t, ErrNegativeBalance, err)
}

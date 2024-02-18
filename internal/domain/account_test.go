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

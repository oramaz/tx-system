package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)

	tx, err := testQueries.CreateTransfer(
		context.Background(),
		CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		},
	)

	require.NoError(t, err)
	require.NotEmpty(t, tx)

	require.NotZero(t, tx.ID)
	require.Equal(t, tx.FromAccountID, account1.ID)
	require.Equal(t, tx.ToAccountID, account2.ID)
	require.Equal(t, tx.Amount, amount)
	require.NotZero(t, tx.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)

	tx1, err := testQueries.CreateTransfer(
		context.Background(),
		CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		},
	)
	require.NoError(t, err)

	tx2, err := testQueries.GetTransfer(context.Background(), tx1.ID)
	require.NoError(t, err)

	require.Equal(t, tx1.ID, tx2.ID)
	require.Equal(t, tx1.FromAccountID, tx2.FromAccountID)
	require.Equal(t, tx1.ToAccountID, tx2.ToAccountID)
	require.Equal(t, tx1.Amount, tx2.Amount)
	require.WithinDuration(t, tx1.CreatedAt, tx2.CreatedAt, time.Second)
}

package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	amount := int64(10)

	entry, err := testQueries.CreateEntry(
		context.Background(),
		CreateEntryParams{
			AccountID: account1.ID,
			Amount:    amount,
		},
	)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.Equal(t, entry.AccountID, account1.ID)
	require.Equal(t, entry.Amount, amount)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	amount := int64(10)

	entry1, err := testQueries.CreateEntry(
		context.Background(),
		CreateEntryParams{
			AccountID: account1.ID,
			Amount:    amount,
		},
	)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry2.CreatedAt, entry2.CreatedAt, time.Second)
}

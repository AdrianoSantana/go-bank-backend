package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/AdrianoSantana/go-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func createParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    sql.NullString{String: util.RandomString(12), Valid: true},
		Balance:  sql.NullInt64{Int64: int64(util.RandomInt(1, 1000)), Valid: true},
		Currency: sql.NullString{String: "USD", Valid: true},
	}
}

func Test_account(t *testing.T) {
	arg := createParams()

	account, err := testQueries.CreateAccount(context.Background(), arg)
	verifyErrorAndEmptyResult(t, err, account)

	require.Equal(t, arg.Owner.String, account.Owner.String)
	require.Equal(t, arg.Balance.Int64, account.Balance.Int64)
	require.Equal(t, arg.Currency.String, account.Currency.String)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	account := createMockAccount(t)

	result, err := testQueries.GetAccount(context.Background(), account.ID)
	verifyErrorAndEmptyResult(t, err, result)

	require.Equal(t, result.ID, account.ID)
	require.Equal(t, result.Owner.String, account.Owner.String)
	require.Equal(t, result.Balance.Int64, account.Balance.Int64)
	require.Equal(t, result.Currency.String, account.Currency.String)
}

func TestUpdateAccou(t *testing.T) {
	account := createMockAccount(t)
	updateParams := UpdateAccountParams{
		ID:      account.ID,
		Balance: sql.NullInt64{Int64: int64(util.RandomInt(1, 1000)), Valid: true},
	}
	result, err := testQueries.UpdateAccount(context.Background(), updateParams)

	verifyErrorAndEmptyResult(t, err, result)
	require.Equal(t, updateParams.Balance.Int64, result.Balance.Int64)
}

func TestDeleteAccount(t *testing.T) {
	account := createMockAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	var accounts []Account
	for i := 0; i < 10; i++ {
		accounts = append(accounts, createMockAccount(t))
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	result, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, result, 5)

	for _, account := range result {
		require.NotEmpty(t, account)
	}
}

func createMockAccount(t *testing.T) Account {
	arg := createParams()

	account, err := testQueries.CreateAccount(context.Background(), arg)
	verifyErrorAndEmptyResult(t, err, account)

	return account
}

func verifyErrorAndEmptyResult(t *testing.T, err error, result Account) {
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

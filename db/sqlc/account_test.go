package db

import (
	"context"
	"github/atulkumar0001/Bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T)(Accounts,CreateAccountParams){
	user,_ := CreateRandomUser(t)
	arg := CreateAccountParams{
		Owner: user.Username, 
		Balance: util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}

	account,err := testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	return account,arg
	
}

func TestCreateAccount(t *testing.T){
	account,arg := CreateRandomAccount(t)

	require.NotEmpty(t,account)
	require.Equal(t,arg.Owner,account.Owner)
	require.Equal(t,arg.Balance,account.Balance)
	require.Equal(t,arg.Currency,account.Currency)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)
}

func TestGetAccount(t *testing.T){
	account1,_ := CreateRandomAccount(t)
	account2,err := testQueries.GetAccount(context.Background(),account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t,account2)
	require.Equal(t,account1.Owner,account2.Owner)
	require.Equal(t,account1.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)

}

func TestUpdateAccount(t *testing.T){
	account1,_ := CreateRandomAccount(t)
	account2,err := testQueries.GetAccount(context.Background(),account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t,account2)
	require.Equal(t,account1.Owner,account2.Owner)
	require.Equal(t,account1.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)

	arg := UpdateAccountParams{
		ID: account2.ID,
		Balance: util.RandomAmount(),
	}

	account3,err := testQueries.UpdateAccount(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t,account3)
	require.Equal(t,account3.Balance,arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1,_ := CreateRandomAccount(t)
	id,err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotZero(t,id)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Accounts
	for i := 0; i < 10; i++ {
		lastAccount,_ = CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
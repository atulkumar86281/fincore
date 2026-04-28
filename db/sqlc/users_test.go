package db

import (
	"context"
	"github/atulkumar0001/Bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)


func CreateRandomUser(t *testing.T)(Users,CreateUserParams){
	hashedPass, err := util.HashedPassword(util.RandomString(10))
	require.NoError(t,err)

	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: hashedPass,
		FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
	}

	account,err := testQueries.CreateUser(context.Background(),arg)
	require.NoError(t,err)
	return account,arg
	
}

func TestCreateUser(t *testing.T){
	account,arg := CreateRandomUser(t)

	require.NotEmpty(t,account)
	require.Equal(t,arg.Username,account.Username)
	require.Equal(t,arg.HashedPassword,account.HashedPassword)
	require.Equal(t,arg.FullName,account.FullName)
	require.Equal(t,arg.Email,account.Email)


	require.NotZero(t,account.CreatedAt)
	require.True(t,account.PasswordLastUpdated.IsZero())
}

func TestGetUser(t *testing.T){
	account1,_ := CreateRandomUser(t)
	account2,err := testQueries.GetUser(context.Background(),account1.Username)

	require.NoError(t,err)
	require.NotEmpty(t,account2)
	require.Equal(t,account1.Username,account2.Username)
	require.Equal(t,account1.HashedPassword,account2.HashedPassword)
	require.Equal(t,account1.FullName,account2.FullName)
	require.Equal(t,account1.Email,account2.Email)
}
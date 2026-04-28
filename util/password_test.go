package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)


func TestPassword(t *testing.T){
	password := RandomString(16)

	pass1, err := HashedPassword(password)
	require.NoError(t,err)
	require.NotEmpty(t,pass1)

	err = CheckPassword(password,pass1);
	require.NoError(t,err)

	wrongPass := RandomString(16);
	err = CheckPassword(wrongPass,pass1)
	require.EqualError(t,err,bcrypt.ErrMismatchedHashAndPassword.Error())

	pass2, err := HashedPassword(password)
	require.NoError(t,err)
	require.NotEmpty(t,pass2)
	require.NotEqual(t,pass1,pass2)
}
package token

import (
	"github/atulkumar0001/Bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func TestJwtMaker(t *testing.T){
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t,err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username,duration)
	require.NotEmpty(t,payload)
	require.NoError(t,err)
	require.NotEmpty(t,token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t,err)
	require.NotEmpty(t,payload)


	require.NotZero(t,payload.ID)
	require.Equal(t,username,payload.Username)
	require.WithinDuration(t,issuedAt,payload.IssuedAt,time.Second)
	require.WithinDuration(t,expiredAt,payload.ExpiredAt,time.Second)

}

func TestExpiredToken(t *testing.T){
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t,err)

	username := util.RandomOwner()
	duration := -time.Minute

	token, payload, err := maker.CreateToken(username,duration)
	require.NotEmpty(t,payload)
	require.NoError(t,err)
	require.NotEmpty(t,token)

	payload, err = maker.VerifyToken(token)
	require.Error(t,err)
	require.Nil(t,payload)
	require.EqualError(t,err, ErrExpiredToken.Error())

}


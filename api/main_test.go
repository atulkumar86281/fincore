package api

import (
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/util"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server{
	config := util.Config{
		TokenAssymetricKey: util.RandomString(32),
		AcccessTokenDuration: time.Minute,
	}

	sever, err := NewServer(config,store)
	require.NoError(t,err)
	return sever
}


func TestMain(m *testing.M){
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
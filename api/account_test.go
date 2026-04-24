package api

import (
	"fmt"
	mockDb "github/atulkumar0001/Bank/db/mock"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)


func TestGetAccountApi(t *testing.T){
	acc := randomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockDb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().GetAccount(gomock.Any(),gomock.Eq(acc.ID)).Times(1).Return(acc,nil)


	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d",acc.ID)
	request, err := http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)

	server.router.ServeHTTP(recorder,request)

	//check response in recorder
	require.Equal(t,http.StatusOK,recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(1,1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
}
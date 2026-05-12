package api

import (
	"database/sql"
	"errors"
	"fmt"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountId    int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountId    int64 `json:"to_account_id" binding:"required,min=1"`
	Amount    int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,currency"`


}
func (server *Server) createTransfer(ctx *gin.Context){

	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	fromAccount, valid := server.validAccount(ctx,req.FromAccountId,req.Currency)
	if !valid{
		return
	}

	if fromAccount.Owner != authPayload.Username{
		err := errors.New("From Account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}


	_,valid = server.validAccount(ctx,req.ToAccountId,req.Currency)
	if !valid{
		return
	}

	arg := db.TransferTxParams{
		FromAccountId: req.FromAccountId,
		ToAccountId: req.ToAccountId,
		Amount: req.Amount,
	}

	result,err := server.store.TransferTx(ctx,arg)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Accounts,bool){
	account, err := server.store.GetAccount(ctx,accountID)

	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return account,false
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return account,false
	}
	if account.Currency != currency{
		err = fmt.Errorf("Account with Id : %d is having currency mismatch, expected : %s , result: %s",account.ID,currency,account.Currency)
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return account,false
	}
	return account,true;
}
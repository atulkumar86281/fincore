package api

import (
	"database/sql"
	"fmt"
	db "github/atulkumar0001/Bank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}
func (server *Server) createAccount(ctx *gin.Context){

	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account,err := server.store.CreateAccount(ctx,arg)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,account)
}

type GetAccountParam struct {
	ID  int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){
	var req GetAccountParam;
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	acc,err := server.store.GetAccount(ctx,req.ID)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,acc)
}

type ListAccountParams struct{
	PageId int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context){
	var req ListAccountParams;
	if err := ctx.ShouldBindQuery(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListAccountParams{
		Limit: req.PageSize,
		Offset: (req.PageId-1) * req.PageSize,
	}
	acc,err := server.store.ListAccount(ctx,arg)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,acc)
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,acc)
}

type DeleteAccountParam struct {
	ID  int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteAccount(ctx *gin.Context){
	var req DeleteAccountParam;
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	_, err := server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type UpdateBalanceParam struct {
	ID  int64 `json:"id" binding:"required,min=1"`
	Balance int64 `json:"balance" binding:"required,min=1"`
}

func (server *Server) UpdateAccount(ctx *gin.Context){
	var req UpdateBalanceParam;
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	fmt.Println("REQ ID:", req.ID)

	arg := db.UpdateAccountParams{
		ID: req.ID,
		Balance: req.Balance,
	}

	acc, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,acc)
}
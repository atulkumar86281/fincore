package api

import (
	"errors"
	"fmt"
	db "github/atulkumar0001/Bank/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type reNewRequest struct {
	RefreshToken    string `json:"refresh_token" binding:"required"`
}

type reNewResponse struct {
	AccessToken string `json:"access_token"`
	AcessTokenExpires time.Time `json:"access_token_expires_at"`
}

func (server *Server) reNewAccessToken(ctx *gin.Context){
	var req reNewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	tokenPayload,err := server.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil{
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}


	session, err := server.store.GetSession(ctx,tokenPayload.ID)

	if err != nil{
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}


	if session.IsBlocked {
		err := fmt.Errorf("Session is blocked")
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	if req.RefreshToken != session.RefreshToken{
		err := fmt.Errorf("Refresh Token Mismatch")
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	if tokenPayload.Username != session.Username{
		err := fmt.Errorf("user mismatch")
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiresAt){
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(tokenPayload.Username,server.config.AcccessTokenDuration)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	res := reNewResponse{
		AccessToken: accessToken,
		AcessTokenExpires: accessTokenPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK,res)
}
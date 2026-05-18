package api

import (
	"errors"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	UserName    string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type userResponse struct{
	Username            string    `json:"username"`
	FullName            string    `json:"full_name"`
	Email               string    `json:"email"`
	PasswordLastUpdated time.Time `json:"password_last_updated"`
	CreatedAt           time.Time `json:"created_at"`
}

func newUserResponse(user db.Users) userResponse{
	return userResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordLastUpdated: user.PasswordLastUpdated,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context){

	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	hashedPass, err := util.HashedPassword(req.Password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.UserName,
		HashedPassword: hashedPass,
		FullName: req.FullName,
		Email: req.Email,
	}

	account,err := server.store.CreateUser(ctx,arg)

	if err != nil{
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden,errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	userArg := newUserResponse(account)

	ctx.JSON(http.StatusOK,userArg)
}

type loginUserRequest struct {
	UserName    string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionId uuid.UUID `json:"session_id"`
	RefreshToken string `json:"refresh_token"`
	RefreshTokenExpires time.Time `json:"refresh_token_expires_at"`
	AccessToken string `json:"access_token"`
	AcessTokenExpires time.Time `json:"access_token_expires_at"`
	User userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context){
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx,req.UserName)

	if err != nil{
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password,user.HashedPassword)

	if err != nil{
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(req.UserName,server.config.AcccessTokenDuration)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(req.UserName,server.config.RefreshTokenDuration)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx,db.CreateSessionParams{
		ID: refreshTokenPayload.ID,
		Username: req.UserName,
		RefreshToken: refreshToken,
		UserAgent: ctx.Request.UserAgent(),
		ClientIp: ctx.ClientIP(),
		IsBlocked: false,
		ExpiresAt: refreshTokenPayload.ExpiredAt,
	})

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}


	res := loginUserResponse{
		SessionId: session.ID,
		RefreshToken: refreshToken,
		RefreshTokenExpires: refreshTokenPayload.ExpiredAt,
		AccessToken: accessToken,
		AcessTokenExpires: accessTokenPayload.ExpiredAt,
		User: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK,res)
}
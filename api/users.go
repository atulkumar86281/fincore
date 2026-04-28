package api

import (
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUserRequest struct {
	UserName    string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreatedUserResponse struct{
	Username            string    `json:"username"`
	FullName            string    `json:"full_name"`
	Email               string    `json:"email"`
	PasswordLastUpdated time.Time `json:"password_last_updated"`
	CreatedAt           time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context){

	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	hashedPass, err := util.HashedPassword(req.Password)
	if err != nil{
		if pqErr, ok := err.(*pq.Error); ok{
			switch pqErr.Code.Name(){
			case "unique_violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			}
		}
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
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	userArg := CreatedUserResponse{
		Username: account.Username,
		FullName: account.FullName,
		Email: account.Email,
		PasswordLastUpdated: account.PasswordLastUpdated,
		CreatedAt: account.CreatedAt,
	}

	ctx.JSON(http.StatusOK,userArg)
}
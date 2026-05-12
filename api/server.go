package api

import (
	"fmt"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/token"
	"github/atulkumar0001/Bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct{
	config util.Config
	tokenMaker token.Maker
	store db.Store
	router *gin.Engine
}

func NewServer(config util.Config,store db.Store) (*Server,error){
	tokenMaker, err := token.NewPasetoMaker(config.TokenAssymetricKey)
	if err != nil{
		return nil, fmt.Errorf("Could not initialize the token maker: %w",err)
	}
	server := &Server{
		store: store,
		tokenMaker: tokenMaker,
		config: config,
	}

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validateCurrency)
	}

	server.setupRouter()

	return server,nil;
}

func (server *Server) Start(address string)error{
	return server.router.Run(address);
}

func errorResponse(err error) gin.H{
	return gin.H{"error":err.Error()}
}

func (server *Server) setupRouter(){
	router := gin.Default()

	authGroup := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// add routes to the router
	authGroup.POST("/accounts",server.createAccount)
	authGroup.GET("/accounts/:id",server.getAccount)
	authGroup.GET("/accounts",server.listAccount)
	authGroup.DELETE("/accounts/:id",server.DeleteAccount)
	authGroup.PUT("/accounts/",server.UpdateAccount)
	authGroup.POST("/transfers",server.createTransfer)


	router.POST("/users",server.createUser)
	router.POST("/users/login",server.loginUser)

	server.router = router
}
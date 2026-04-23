package api

import (
	db "github/atulkumar0001/Bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct{
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	// add routes to the router
	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)
	router.DELETE("/accounts/:id",server.DeleteAccount)
	router.PUT("/accounts/",server.UpdateAccount)



	server.router = router
	return server;
}

func (server *Server) Start(address string)error{
	return server.router.Run(address);
}

func errorResponse(err error) gin.H{
	return gin.H{"error":err.Error()}
}
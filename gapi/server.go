package gapi

import (
	"fmt"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/pb"
	"github/atulkumar0001/Bank/token"
	"github/atulkumar0001/Bank/util"

	"github.com/gin-gonic/gin"
)


type Server struct{
	pb.UnimplementedBankServer
	config util.Config
	tokenMaker token.Maker
	store db.Store
	router *gin.Engine
}

// Creates a new grpc server 
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


	return server,nil;
}
package main

import (
	"database/sql"
	"github/atulkumar0001/Bank/api"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/gapi"
	"github/atulkumar0001/Bank/pb"
	"github/atulkumar0001/Bank/util"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main(){
	config,err := util.LoadConfig(".")

	if err != nil{
		log.Fatal("couldn't load the config file")
	}

	
	conn,err := sql.Open(config.DBDriver,config.DBSource)

	if err != nil{
		log.Fatal("Cannot Connect to the database: ",err)
	}

	store := db.NewStore(conn)

	// running grpc server, switch this to run normal http gin server
	runGrpcServer(config,store)
	
}
func runGrpcServer(config util.Config, store db.Store){
	server, err := gapi.NewServer(config,store)

	if err != nil{
		log.Fatal("Something went wrong with token maker initialization")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankServer(grpcServer,server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp",config.GrpcServerAddress)

	if err != nil{
		log.Fatal("Cannot Create Listener")
	}

	log.Printf("Grpc Server Started at: %s",listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil{
		log.Fatal("Cannot start grpc server")
	}

}

func runGinServer(config util.Config, store db.Store){
	server, err := api.NewServer(config,store)

	if err != nil{
		log.Fatal("Something went wrong with token maker initialization")
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil{
		log.Fatal("Cannot start the server",err)
	}
}
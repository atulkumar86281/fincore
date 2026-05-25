package main

import (
	"context"
	"database/sql"
	"github/atulkumar0001/Bank/api"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/gapi"
	"github/atulkumar0001/Bank/pb"
	"github/atulkumar0001/Bank/util"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)


func main(){
	config,err := util.LoadConfig(".")

	if err != nil{
		log.Fatal("couldn't load the config file: ",err)
	}

	
	conn,err := sql.Open(config.DBDriver,config.DBSource)

	if err != nil{
		log.Fatal("Cannot Connect to the database: ",err)
	}

	store := db.NewStore(conn)

	// running both server
	go runGatewayServer(config,store)
	// running grpc server, switch this to run normal http gin server
	runGrpcServer(config,store)
	
}
func runGrpcServer(config util.Config, store db.Store){
	server, err := gapi.NewServer(config,store)

	if err != nil{
		log.Fatal("Something went wrong with token maker initialization: ",err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankServer(grpcServer,server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp",config.GrpcServerAddress)

	if err != nil{
		log.Fatal("Cannot Create Listener: ",err)
	}

	log.Printf("Grpc Server Started at: %s",listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil{
		log.Fatal("Cannot start grpc server: ",err)
	}

}

func runGinServer(config util.Config, store db.Store){
	server, err := api.NewServer(config,store)

	if err != nil{
		log.Fatal("Something went wrong with token maker initialization: ", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil{
		log.Fatal("Cannot start the server: ",err)
	}
}


func runGatewayServer(config util.Config, store db.Store){
	server, err := gapi.NewServer(config,store)

	if err != nil{
		log.Fatal("Something went wrong with token maker initialization: ", err)
	}

	json := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(json)

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterBankHandlerServer(ctx,grpcMux,server)

	if err != nil{
		log.Fatal("Cannot register handler server: ",err)
	}

	mux := http.NewServeMux()
	mux.Handle("/",grpcMux)


	listener, err := net.Listen("tcp",config.HttpServerAddress)

	if err != nil{
		log.Fatal("Cannot Create Listener: ",err)
	}

	log.Printf("Http Server Started at: %s",listener.Addr().String())

	err = http.Serve(listener,mux)
	if err != nil{
		log.Fatal("Cannot start grpc server: ",err)
	}

}
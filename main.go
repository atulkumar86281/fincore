package main

import (
	"database/sql"
	"github/atulkumar0001/Bank/api"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/util"
	"log"

	_ "github.com/lib/pq"
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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil{
		log.Fatal("Cannot start the server",err)
	}
}
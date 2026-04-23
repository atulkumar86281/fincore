package db

import (
	"database/sql"
	"github/atulkumar0001/Bank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB


func TestMain(m *testing.M){
	config,err := util.LoadConfig("../../")

	if err != nil{
		log.Fatal("couldn't load the config file")
	}

	testDb,err = sql.Open(config.DBDriver,config.DBSource)

	if err != nil{
		log.Fatal("Cannot Connect to the database: ",err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run());
}
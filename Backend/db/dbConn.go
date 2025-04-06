package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB


func InitDB(){
	connStr := "host=localhost port=5432 password=soham dbname=TrendMind sslmode=disable"
	
	db, err := sql.Open("postgres", connStr)
	
	if err != nil{
		log.Fatal(err)
	}
	
	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}
}

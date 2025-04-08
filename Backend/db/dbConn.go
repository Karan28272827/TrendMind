package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=trendmind sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))

	fmt.Println(os.Getenv("DB_USER"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Connected to postgres")
	DB = db
}

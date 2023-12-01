package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	db "github.com/darionatias-dev/todo_golang/api/db/sqlc"
)

var dbQueries *db.Queries

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dbDriver := os.Getenv("DATABASE_DRIVER")
	dbUrl := os.Getenv("DATABASE_URL")

	dbcon, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbcon.Close()

	if err := dbcon.Ping(); err != nil {
		log.Fatal(dbcon)
	}

	dbQueries = db.New(dbcon)
}

func main() {}

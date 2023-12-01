package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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

func main() {
	app := gin.Default()

	if err := app.Run("localhost:3001"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	db "github.com/dariomatias-dev/todo_golang/api/db/sqlc"
	"github.com/dariomatias-dev/todo_golang/api/routes"
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
	defer dbcon.Close()

	if err := dbcon.Ping(); err != nil {
		log.Fatal(err)
	}

	dbQueries = db.New(dbcon)
}

func main() {
	app := gin.Default()

	routes.AppRoutes(app, dbQueries)

	if err := app.Run("localhost:3001"); err != nil {
		log.Fatal(err)
	}
}

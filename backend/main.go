package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/v420v/qrmarkapi/api"
)

func connectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, dbName))
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	r := api.NewRouter(db)

	_ = godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Printf("server start QrmarkAPI at port %s\n", ":"+PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}

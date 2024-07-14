package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"math/rand"

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

func AddData(db *sql.DB) {
	time.Sleep(10)
	id := 700
	user_id := 1
	school_id := 7026

	for i := 0; i < 100000; i++ {
		company_id := rand.Intn(29) + 1
		points := 3
		createdAt := time.Now()
		_, err := db.Exec(`insert into
		qrmarks (qrmark_id, user_id, school_id, company_id, points, created_at)
		values (?, ?, ?, ?, ?, ?);`, id+i, user_id, school_id, company_id, points, createdAt)
		if err != nil {
			panic(err)
		}
	}
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

	//AddData(db)

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}

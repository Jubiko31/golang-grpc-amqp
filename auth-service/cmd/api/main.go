package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

const webPort = "8001"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("🔐 Listening Authentication Service...")

	env_err := godotenv.Load(".env")
	if env_err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn := connectToDB()
	if conn == nil {
		log.Panic("Failed to connect to PostgreSQL.")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("Failed to connect to PostgreSQL.")
			counts++
		} else {
			log.Println("🐘 Connected to PostgreSQL")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

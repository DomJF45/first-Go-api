package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt"`
	LastLogin    time.Time `json:"lastLogin"`
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// load sql config
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// get db handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// get ping
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// load and init controller
	albumController := newAlbumController()
	controllers := []ControllerInterface{
		albumController,
	}
	app := newApp(controllers, "localhost:8080")
	app.run()
}

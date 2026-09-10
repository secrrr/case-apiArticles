package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
    "time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    connStr := os.Getenv("DB_CONNECTION")
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Gagal inisialisasi db:", err)
    }
    
    if err := db.Ping(); err != nil {
        log.Fatal("Database tidak merespon:", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Koneksi database berhasil!")
    
    return db
}
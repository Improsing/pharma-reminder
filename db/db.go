package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//const dsn = "user=postgres password=123 dbname=medscheduler sslmode=disable"

//func Connect() (*sql.DB, error) {
//	db, err := sql.Open("postgres", dsn)
//		if err != nil {
//			return nil, err
//		}
//		return db, db.Ping()
//	}

func getDBConfig() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func Connect() (*sql.DB, error) {
	dsn := getDBConfig()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func InitSchema(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS schedules (
		id SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL,
		medicine_name TEXT NOT NULL,
		frequency INT NOT NULL,
		duration INT,
		start_time TIMESTAMP NOT NULL
	);
	ALTER TABLE schedules ADD COLUMN IF NOT EXISTS notes TEXT;
	`
_, err := db.Exec(query)
if err != nil {
	log.Fatal("Failed to initialize schema:", err)
	}
}


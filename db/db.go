package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const dsn = "user=postgres password=pharma dbname=medscheduler sslmode=disable"

func Connect() (*sql.DB, error) {
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
);`
_, err := db.Exec(query)
if err != nil {
	log.Fatal("Failed to initialize schema:", err)
	}
}


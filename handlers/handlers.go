package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Improsing/pharma-reminder/models"
)


func CreateSchedule(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var schedule models.Schedule
		if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO schedules (user_id, medicine_name, frequency, duration, start_time)
				  VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err := db.QueryRow(query, schedule.UserID, schedule.MedicineName, schedule.Frequency, schedule.Duration, schedule.StartTime).Scan(&schedule.ID)
		if err != nil {
			log.Println("Error inserting schedule:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": schedule.ID})
	}
}

func GetSchedules(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT id FROM schedules WHERE user_id = $1", userID)
		if err != nil {
			log.Println("Error fetching schedules", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var scheduleIDs []int
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				http.Error(w, "Error scanning database results", http.StatusInternalServerError)
				return
			}
			scheduleIDs = append(scheduleIDs, id)
		}

		json.NewEncoder(w).Encode(scheduleIDs)
	}
}

func GetSchedule(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		scheduleID := r.URL.Query().Get("schedule_id") 

		if userID == "" || scheduleID == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}

		var schedule models.Schedule
		query := `SELECT id, user_id, medicine_name, frequency, duration, start_time FROM schedules WHERE id = $1 AND user_id = $2`
		err := db.QueryRow(query, scheduleID, userID).Scan(&scheduleID, &schedule.UserID, &schedule.MedicineName, &schedule.Frequency, schedule.Duration, schedule.StartTime)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Schedule not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(schedule)
	}
}
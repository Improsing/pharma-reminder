package models

import "time"

type Schedule struct {
	ID            int        `json:"id"`
	UserID        string     `json:"user_id"`
	MedicineName  string     `json:"medicine_name"`
	Frequency     int        `json:"frequency"`
	Duration      *int       `json:"duration,omitempty"`
	StartTime     time.Time  `json:"start_time"`
}
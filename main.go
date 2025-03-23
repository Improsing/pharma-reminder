package main

import (
	"log"
	"net/http"

	"github.com/Improsing/pharma-reminder/db"
	"github.com/Improsing/pharma-reminder/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к БД", err)
	}
	defer database.Close()

	db.InitSchema(database)

	r := mux.NewRouter()
	r.HandleFunc("/schedule", handlers.CreateSchedule(database)).Methods("POST")
	r.HandleFunc("/schedule", handlers.GetSchedules(database)).Methods("GET")
	r.HandleFunc("/schedule", handlers.GetSchedule(database)).Methods("GET")

	log.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	
}

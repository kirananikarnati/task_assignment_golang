package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_example/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/segmentio/kafka-go"
)

var db *gorm.DB

func main() {
	// Connect to the database
	var err error
	db, err = gorm.Open("postgres", "host=mahmud.db.elephantsql.com port=5432 user=vkszobhj dbname=vkszobhj password=FUwHIAWaV68Ecg1DmX-sulMeOsLGUWUN sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create tasks table if it doesn't exist
	if !db.HasTable(&models.Task{}) {
		db.CreateTable(&models.Task{})
	}

	// Create a new kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "topic-task",
		Balancer: &kafka.RoundRobin{},
	})
	defer writer.Close()

	// Create HTTP server
	http.HandleFunc("/createtask", func(w http.ResponseWriter, r *http.Request) {
		// Parse input from request
		var taskInput models.Task
		if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Insert task into the database and get the generated ID
		task := &models.Task{Input: taskInput.Input, Status: "not_started"}
		db.Create(task)

		// Send task to kafka topic
		message := kafka.Message{
			Value: []byte(fmt.Sprintf("%d", task.ID)),
		}
		if err := writer.WriteMessages(context.Background(), message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return task ID to the client
		json.NewEncoder(w).Encode(task.ID)
	})

	http.HandleFunc("/retrievetask", func(w http.ResponseWriter, r *http.Request) {
		// Parse task ID from request
		var id int
		if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Find task in the database
		task := &models.Task{}
		if err := db.First(task, id).Error; err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		// Return task to the client
		json.NewEncoder(w).Encode(task)
	})

	http.HandleFunc("/taskresult", func(w http.ResponseWriter, r *http.Request) {
		// Parse task ID from request
		var id int
		if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Find task in the database
		task := &models.Task{}
		if err := db.First(task, id).Error; err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		// Return task result to the client
		json.NewEncoder(w).Encode(task.Output)
	})

	// Start HTTP server
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

}

package main

import (
	"context"
	"fmt"
	"log"
	"rest_example/models"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/segmentio/kafka-go"
)

func main() {
	// setup kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "topic-task",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	// Connect to the PostgreSQL database
	db, err := gorm.Open("postgres", "host=mahmud.db.elephantsql.com port=5432 user=vkszobhj dbname=vkszobhj password=FUwHIAWaV68Ecg1DmX-sulMeOsLGUWUN sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read messages from kafka topic
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		num := string(m.Value)
		taskid, _ := strconv.ParseInt(num, 0, 16)
		var task models.Task
		db.First(&task, taskid)
		time.Sleep(time.Second)
		if task.Status == "not_started" {
			db.Model(&task).Where("id = ?", taskid).Updates(models.Task{Status: "success", Output: strings.ToUpper(task.Input)})
			fmt.Println("Processing successful for the task : ", taskid)
		}
		// else {
		// 	fmt.Println("Already Proccessed.")
		// }
		reader.CommitMessages(context.Background(), m)
	}
}

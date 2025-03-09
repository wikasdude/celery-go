package handler

import (
	"celery-go/internal/queue"
	"celery-go/internal/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type TaskMessage struct {
	TaskID string `json:"task_id"`
	Data   string `json:"data"` // Actual task data
}

func SubmitTask(q *queue.RabbitMQ, w http.ResponseWriter, r *http.Request) {
	storage.InitDB()
	storage.InitElasticsearch()
	taskID := uuid.New().String() // Generate a unique task ID
	taskData := "Process This Task"

	taskMsg := TaskMessage{
		TaskID: taskID,
		Data:   taskData,
	}
	msgBody, err := json.Marshal(taskMsg)
	if err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}
	//taskID := uuid.New().String()
	err = storage.SaveTask(taskID, "pending")
	if err != nil {
		storage.LogTaskToElasticsearch(taskID, "failed", "Failed to submit task to queue")
		http.Error(w, "Failed to save task", http.StatusInternalServerError)
		return
	}
	//task := "Process This Task"
	err = q.Publish("tasks", string(msgBody))
	if err != nil {
		storage.UpdateTaskStatus(taskID, "failed", "")
		http.Error(w, "Failed to submit task", http.StatusInternalServerError)
		return
	}
	// Set the content-type header to application/json for a structured response
	w.Header().Set("Content-Type", "application/json")

	// Respond to the client with a success message
	w.WriteHeader(http.StatusOK)
	storage.UpdateTaskStatus(taskID, "submitted", "")
	storage.LogTaskToElasticsearch(taskID, "submitted", "Task successfully submitted")
	log.Println("log successfully ")
	fmt.Fprintf(w, "Task submitted successfully")
}

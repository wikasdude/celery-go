package consumer

import (
	"celery-go/internal/storage"
	"log"
)

// TaskPayload defines the structure of a task message
type TaskPayload struct {
	TaskID string `json:"task_id"`
	Task   string `json:"task"`
}

// ProcessTask processes the task and updates the database
func ProcessTask(task TaskPayload) error {
	storage.InitDB()
	log.Printf("Processing task: %s, ID: %s\n", task.Task, task.TaskID)

	// Update status to "in_progress"
	err := storage.UpdateTaskStatus(task.TaskID, "in_progress", "")
	if err != nil {
		log.Printf("Failed to update task %s to in_progress: %v\n", task.TaskID, err)
		return err
	}

	// Simulate processing
	result := "Task completed successfully"

	// Update status to "completed"
	err = storage.UpdateTaskStatus(task.TaskID, "completed", result)
	if err != nil {
		log.Printf("Failed to update task %s to completed: %v\n", task.TaskID, err)
		return err
	}

	log.Printf("Task ID %s marked as completed\n", task.TaskID)
	return nil
}

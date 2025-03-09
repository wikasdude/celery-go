package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := "postgres://myuser:mypassword@postgres:5432/celery_tasks?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		task_id TEXT UNIQUE NOT NULL,
		status TEXT NOT NULL,
		result TEXT
	)`)
	if err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully")
}

func SaveTask(taskID, status string) error {
	query := `INSERT INTO tasks (task_id, status) VALUES ($1, $2) ON CONFLICT (task_id) DO NOTHING`
	_, err := db.Exec(query, taskID, status)
	if err != nil {
		return fmt.Errorf("failed to save task: %v", err)
	}
	return nil
}

// UpdateTaskStatus updates the status of an existing task
func UpdateTaskStatus(taskID, status, result string) error {
	query := `UPDATE tasks SET status = $1, result = $2 WHERE task_id = $3`
	_, err := db.Exec(query, status, result, taskID)
	if err != nil {
		return fmt.Errorf("failed to update task status: %v", err)
	}
	return nil
}

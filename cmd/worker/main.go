package main

import (
	"celery-go/cmd/worker/consumer"
	"celery-go/internal/queue"
	"encoding/json"
	"log"
)

func main() {
	q, err := queue.NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	msgs, err := q.Consume("tasks")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Worker is waiting for tasks...")
	for msg := range msgs {
		var task consumer.TaskPayload
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Println("Failed to parse task:", err)
			continue
		}
		log.Printf("Processing task: %s, ID: %s\n", task.Task, task.TaskID)
		//processTask(string(msg.Body))
		err = consumer.ProcessTask(task)
		if err != nil {
			log.Printf("Failed to process task %s: %v\n", task.TaskID, err)
			msg.Nack(false, true) // Requeue if failed
			continue
		}

		// Acknowledge message after successful processing
		msg.Ack(false)
		//msg.Ack(false)
	}
}

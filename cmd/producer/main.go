package main

import (
	"celery-go/cmd/producer/auth"
	"celery-go/cmd/producer/handler"
	"celery-go/cmd/producer/middleware"
	"celery-go/internal/queue"
	"log"
	"net/http"
)

var q *queue.RabbitMQ

func submitTaskHandler(w http.ResponseWriter, r *http.Request, q *queue.RabbitMQ) {
	// Extract JWT claims from the request
	claims, err := middleware.VerifyJWT(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Log the claims to see if the JWT is parsed correctly
	log.Printf("JWT Claims: %+v", claims)

	// Role-based access control
	if claims.Role != "admin" && claims.Role != "client" {
		http.Error(w, "Forbidden: You don't have permission", http.StatusForbidden)
		return
	}

	// Handle task submission
	handler.SubmitTask(q, w, r)

}

func main() {
	q, err := queue.NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()
	token, _ := auth.GenerateJWT("vikas_sharma", "client")
	log.Println("JWT:", token)

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		submitTaskHandler(w, r, q)
	})

	log.Println("Producer running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// Initialize Elasticsearch client
var esClient *elasticsearch.Client

func InitElasticsearch() {
	var err error
	esClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200", // Elasticsearch URL
		},
	})

	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
}

// LogTaskToElasticsearch sends task logs to Elasticsearch
func LogTaskToElasticsearch(taskID, status, message string) error {
	// Prepare the log entry
	logEntry := map[string]interface{}{
		"task_id":   taskID,
		"status":    status,
		"message":   message,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	// Marshal logEntry to JSON
	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		return fmt.Errorf("error marshaling log entry: %s", err)
	}

	// Index the log into Elasticsearch
	resp, err := esClient.Index(
		"task_logs",                           // Index name
		bytes.NewReader(logJSON),              // Document content (log entry)
		esClient.Index.WithDocumentID(taskID), // Document ID
		esClient.Index.WithRefresh("true"),    // Refresh index after indexing
	)

	if err != nil {
		return fmt.Errorf("error indexing document: %s", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.IsError() {
		return fmt.Errorf("error indexing document: %s", resp.String())
	}
	return nil
}

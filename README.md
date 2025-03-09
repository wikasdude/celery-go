# Celery-Go

Celery-Go is a distributed task queue system built in Golang, inspired by Celery (Python). It enables asynchronous task execution using message brokers like Redis and RabbitMQ.

## Features
- Asynchronous task execution
- Support for multiple message brokers (Redis, RabbitMQ)
- Worker process management
- Task persistence
- Kubernetes support (upcoming)
- Monitoring via Grafana

## Project Structure
```
celery-go/
├── cmd/
│   ├── producer/
│   │   ├── main.go
│   │   ├── handler.go
│   │   ├── rabbitmq.go
│   ├── worker/
│   │   ├── main.go
│   │   ├── consumer.go
├── config/
├── internal/
│   ├── queue/
│   ├── storage/
├── api/
├── pkg/
├── docker-compose.yml
├── README.md
```

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/celery-go.git
   cd celery-go
   ```
2. Build the project:
   ```sh
   go build ./...
   ```
3. Run with Docker:
   ```sh
   docker-compose up --build
   ```

## Usage
### Start a Producer
```sh
cd cmd/producer
go run main.go
```

### Start a Worker
```sh
cd cmd/worker
go run main.go
```

## Monitoring with Grafana
- Prometheus metrics are exposed at `/metrics`.
- Use Grafana to visualize worker performance.

## Roadmap
- [ ] Add Kubernetes support
- [ ] Implement a REST API for task management
- [ ] Enhance logging and monitoring

## License
MIT License


FROM golang:1.22

WORKDIR /app

# Copy only go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application
# Copy only the worker files
COPY . .

# Set working directory to worker
WORKDIR /app/cmd/worker

# Build the worker
RUN go build -o worker

# Run the worker
CMD ["./worker"]

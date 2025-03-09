FROM golang:1.22

WORKDIR /app

# Copy only go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies early
RUN go mod tidy
# Copy only the producer files
COPY . .

# Set working directory to producer
WORKDIR /app/cmd/producer

# Build the application
RUN go build -o producer

# Run the producer
CMD ["./producer"]

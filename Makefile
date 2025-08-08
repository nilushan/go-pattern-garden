# Name of the output binary
BINARY_NAME=patterngarden-app
# Path to the main package
MAIN_PATH=./cmd/patterngarden
# Path to the docker-compose file
COMPOSE_FILE=./deploy/docker-compose.yml

.PHONY: help run build clean docker-up docker-down docker-logs

help:
	@echo "Available commands:"
	@echo "  make run           - Run the Go application locally"
	@echo "  make build         - Build the Go application binary"
	@echo "  make clean         - Remove the built binary"
	@echo "  make docker-up     - Start all services with Docker Compose in detached mode"
	@echo "  make docker-down   - Stop all services and remove volumes"
	@echo "  make docker-logs   - View logs from all running services"

# --- Go Commands ---
run:
	@echo "Running the application..."
	go run $(MAIN_PATH)

build:
	@echo "Building binary..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

clean:
	@echo "Cleaning up..."
	@rm -f bin/$(BINARY_NAME)

# --- Docker Compose Commands ---
docker-up:
	@echo "Starting Docker services..."
	docker-compose -f $(COMPOSE_FILE) up --build -d

docker-down:
	@echo "Stopping Docker services and removing volumes..."
	docker-compose -f $(COMPOSE_FILE) down -v

docker-logs:
	@echo "Tailing logs..."
	docker-compose -f $(COMPOSE_FILE) logs -f

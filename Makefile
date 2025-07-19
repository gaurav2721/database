# Makefile for B+Tree Index Project
# Variables
BINARY_NAME=database
MAIN_FILE=main.go
BUILD_DIR=build
TEST_DIR=tests
DOCKER_IMAGE_NAME=db-bptree
DOCKER_CONTAINER_NAME=db-bptree-container

# Go commands 
GO=go
GOBUILD=$(GO) build
GOTEST=$(GO) test

# Docker commands
DOCKER=docker

.PHONY: build run test docker-build docker-run docker-stop docker-clean

# Build the application
build:
	@echo "Running go mod tidy and downloading dependencies..."
	$(GO) mod tidy
	$(GO) mod download
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run:
	@echo "Running go mod tidy and downloading dependencies..."
	$(GO) mod tidy
	$(GO) mod download
	@echo "Running $(BINARY_NAME)..."
	$(GO) run $(MAIN_FILE)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./$(TEST_DIR)/...

# Docker commands

# Build Docker image
docker-build:
	@echo "Building Docker image $(DOCKER_IMAGE_NAME)..."
	$(DOCKER) build -t $(DOCKER_IMAGE_NAME) .
	@echo "Docker image build complete: $(DOCKER_IMAGE_NAME)"

# Run Docker container
docker-run:
	@echo "Running Docker container $(DOCKER_CONTAINER_NAME)..."
	$(DOCKER) run --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

# Stop Docker container
docker-stop:
	@echo "Stopping Docker container $(DOCKER_CONTAINER_NAME)..."
	$(DOCKER) stop $(DOCKER_CONTAINER_NAME) || true
	$(DOCKER) rm $(DOCKER_CONTAINER_NAME) || true

# Clean Docker resources
docker-clean:
	@echo "Cleaning Docker resources..."
	$(DOCKER) stop $(DOCKER_CONTAINER_NAME) || true
	$(DOCKER) rm $(DOCKER_CONTAINER_NAME) || true
	$(DOCKER) rmi $(DOCKER_IMAGE_NAME) || true

# Build and run Docker container
docker-build-run: docker-build docker-run

# Show Docker logs
docker-logs:
	@echo "Showing logs for container $(DOCKER_CONTAINER_NAME)..."
	$(DOCKER) logs $(DOCKER_CONTAINER_NAME)

# Show Docker container status
docker-status:
	@echo "Docker container status:"
	$(DOCKER) ps -a | grep $(DOCKER_CONTAINER_NAME) || echo "Container not found"
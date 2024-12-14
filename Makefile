# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
API_APP_NAME=tasksapi

# Build the API
build-api:
	$(GOBUILD) -o $(API_APP_NAME) cmd/api/main.go

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(API_APP_NAME)

# Install dependencies
deps:
	$(GOGET) -v ./...

# Run the application
run:
	$(GOBUILD) -o $(API_APP_NAME) -v
	./$(API_APP_NAME)

.PHONY: build test clean deps run build-api

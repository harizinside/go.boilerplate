APP_NAME := Go.Boilerplate

.PHONY: run build clean

run:
	@echo "Running the application..."
	go run cmd/app/main.go

build:
	@echo "Building the binary..."
	go build cmd/app/main.go -o $(APP_NAME) .
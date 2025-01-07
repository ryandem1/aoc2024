# Makefile for aoc2024 (Advent of Code)

APP_NAME = aoc2024
BUILD_DIR = ./bin

.PHONY: all build clean run test

# Build the project
all: build

# Build the Go binary
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./*.go
	chmod +x $(BUILD_DIR)/$(APP_NAME)

# Run the project
run: build
	$(BUILD_DIR)/$(APP_NAME)

# Clean the build
clean:
	rm -rf $(BUILD_DIR)

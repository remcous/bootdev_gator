# Define the name of the executable
BINARY_NAME=gator

# Define the Go source file
SRC=main.go

# Define the Goose command and the migration directory
GOOSE_CMD=goose
MIGRATION_DIR=./sql/schema
DB_CONNECTION="postgres://postgres:postgres@localhost:5432/gator"

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME) $(SRC)

# Clean up the build artifacts and run goose down migration
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@if [ -f $(BINARY_NAME) ]; then \
		rm -f $(BINARY_NAME); \
		echo "Removed executable: $(BINARY_NAME)"; \
	fi
	@echo "Running goose down migration..."
	$(GOOSE_CMD) postgres $(DB_CONNECTION) down -dir $(MIGRATION_DIR)
	@echo "Running goose up migration..."
	$(GOOSE_CMD) postgres $(DB_CONNECTION) up -dir $(MIGRATION_DIR)

# Help target to display available commands
.PHONY: help
help:
	@echo "Makefile Commands:"
	@echo "  make build  - Build the application"
	@echo "  make run    - Run the application"
	@echo "  make clean  - Clean up and run goose down migration"
	@echo "  make help   - Show this help message"
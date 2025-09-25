BINARY_NAME=asana
BUILD_DIR=bin

MAKEFLAGS += --silent

.DEFAULT_GOAL := run

.PHONY: run build tidy clean air test

run:
	@echo "🚀 Running application locally..."
	@go run -buildvcs=false cmd/api/main.go

air:
	@echo "🌪️  Starting Air for hot reloading..."
	@if ! command -v air >/dev/null 2>&1; then \
		echo "⚠️  Air not found. Installing Air..."; \
		go install github.com/air-verse/air@latest; \
	fi
	@air

build: 
	@echo "🔨 Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -buildvcs=false -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api/main.go
	@echo "✓ Build completed!"

tidy:
	@echo "📦 Tidying up dependencies..."
	@go mod tidy
	@echo "✓ Dependencies tidied!"

clean:
	@echo "🧹 Cleaning..."
	@go clean
	@rm -rf $(BUILD_DIR) tmp/ docs/api/
	@echo "✓ Cleanup complete!"

test:
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✓ Tests completed!"

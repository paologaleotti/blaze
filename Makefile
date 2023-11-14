CMD_DIRS := $(wildcard cmd/*)
OPENAPI_DIR := api/openapi.yaml

# Define color codes for better output visibility
ORANGE := \033[0;33m
NC := \033[0m  # No Color

# Build all entrypoints and place artifacts in /bin

build: startlog tidy
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${ORANGE}[blaze]${NC} Building ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir)/main $$dir/main.go; \
	done
	@echo "${ORANGE}[blaze]${NC} Done!"


# Clean build artifacts
clean:
	rm -rf bin

generate:
	@echo "${ORANGE}[blaze]${NC} Generating types from OpenAPI schema..."
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate types -package models $(OPENAPI_DIR)
		
startlog:
	@echo "${ORANGE}[blaze]${NC} Build started"

tidy:
	@echo "${ORANGE}[blaze]${NC} Organizing modules..."
	@go mod tidy

lint:
	@echo "${ORANGE}[blaze]${NC} Linting..."
	@golangci-lint run

.PHONY: build clean generate

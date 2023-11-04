CMD_DIRS := $(wildcard cmd/*)
OPENAPI_DIR := api/openapi.yaml
GENERATED_PATH := pkg/models/generated.go

# Define color codes for better output visibility
ORANGE := \033[0;33m
NC := \033[0m  # No Color

# Build all entrypoints and place artifacts in /bin
build: startlog generate tidy
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${ORANGE}[blaze]${NC} Building ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir) $$dir/main.go; \
	done
	@echo "${ORANGE}[blaze]${NC} Done!"

# Build all Lambda entrypoints and place artifacts in /bin
build-lambda: startlog generate tidy
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${ORANGE}[blaze]${NC} Building Lambda for ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir)_lambda $$dir/lambda/main.go; \
	done
	@echo "${ORANGE}[blaze]${NC} Done!"

# Clean build artifacts
clean:
	rm -rf bin

generate:
	@echo "${ORANGE}[blaze]${NC} Generating types from OpenAPI schema..."
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate types -package models $(OPENAPI_DIR) > $(GENERATED_PATH)	
startlog:
	@echo "${ORANGE}[blaze]${NC} Build started"

tidy:
	@echo "${ORANGE}[blaze]${NC} Organizing modules..."
	@go mod tidy

.PHONY: build clean generate

CMD_DIRS := $(wildcard cmd/*)
OPENAPI_DIR := api/openapi.yaml

# Define color codes for better output visibility
ORANGE := \033[0;33m
NC := \033[0m  # No Color
LABEL := ${ORANGE}[blaze]${NC}

# Build all entrypoints and place artifacts in /bin

build: startlog tidy lint-if-env
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${LABEL} Building ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir)/main $$dir/main.go; \
	done
	@echo "${LABEL} Done!"


# Clean build artifacts
clean:
	rm -rf bin

generate:
	@echo "${LABEL} Generating types from OpenAPI schema..."
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate types -package models $(OPENAPI_DIR)
		
startlog:
	@echo "${LABEL} Build started"

tidy:
	@echo "${LABEL} Organizing modules..."
	@go mod tidy

lint:
	@echo "${LABEL} Linting..."
	@golangci-lint run


lint-if-env:
	@if [ -n "$$BLAZE_ALWAYS_LINT" ]; then \
		echo "${LABEL} Linting..."; \
		golangci-lint run; \
	fi

.PHONY: build clean generate

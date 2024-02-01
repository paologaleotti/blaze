CMD_DIRS := $(wildcard cmd/*)
OPENAPI_DIR := api/openapi.yaml
STATIC_DIR := bin/todoservice/static
CLIENT_DIST_DIR := web/client/dist

# Define color codes for better output visibility
ORANGE := \033[0;33m
NC := \033[0m  # No Color

# Build all entrypoints and place artifacts in /bin
build: startlog build-client tidy lint-if-env
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${ORANGE}[blaze]${NC} Building ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir)/main $$dir/main.go; \
	done
	@echo "${ORANGE}[blaze]${NC} Done!"

# Build all Lambda entrypoints and place artifacts in /bin
build-lambda: startlog tidy
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${ORANGE}[blaze]${NC} Building Lambda for ${ORANGE}$$(basename $$dir)${NC}"; \
		GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/$$(basename $$dir)/bootstrap $$dir/lambda/main.go; \
	done
	@echo "${ORANGE}[blaze]${NC} Done!"

# Build client inside web dir using npm and place artifacts in /bin/static
build-client:
	@echo "${ORANGE}[blaze]${NC} Building client..."
	@cd web/client && npm run build
	@rm -rf $(STATIC_DIR)
	@mkdir -p $(STATIC_DIR)
	@mv -f $(CLIENT_DIST_DIR)/* $(STATIC_DIR)

# Clean build artifacts
clean:
	rm -rf $(CLIENT_DIST_DIR)
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


lint-if-env:
	@if [ -n "$$BLAZE_ALWAYS_LINT" ]; then \
		echo "${ORANGE}[blaze]${NC} Linting..."; \
		golangci-lint run; \
	fi

.PHONY: build clean generate

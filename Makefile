CMD_DIRS := $(wildcard cmd/*)

build: startlog tidy lint
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "Building $(basename $$dir)"; \
		go build -o bin/$$(basename $$dir)/main $$dir/main.go; \
	done
	@echo "--- Done! ---"

build-lambda: startlog tidy
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "Building $(basename $$dir)"; \
		GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/$$(basename $$dir)/bootstrap $$dir/lambda/main.go; \
	done
	@echo "--- Done! ---"

clean:
	rm -rf bin
	
startlog:
	@echo "--- Build started ---"

tidy:
	@echo "Organizing modules..."
	@go mod tidy

lint:
	@echo "Linting..."
	@golangci-lint run

.PHONY: build clean tidy lint

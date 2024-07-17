CMD_DIRS := $(wildcard cmd/*)

# Define color codes for better output visibility
ORANGE := \033[0;33m
NC := \033[0m  # No Color
LABEL := ${ORANGE}[blaze]${NC}

build: startlog tidy lint
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "${LABEL} Building ${ORANGE}$$(basename $$dir)${NC}"; \
		go build -o bin/$$(basename $$dir)/main $$dir/main.go; \
	done
	@echo "${LABEL} Done!"

clean:
	rm -rf bin
	
startlog:
	@echo "${LABEL} Build started"

tidy:
	@echo "${LABEL} Organizing modules..."
	@go mod tidy

lint:
	@echo "${LABEL} Linting..."
	@golangci-lint run

.PHONY: build clean
